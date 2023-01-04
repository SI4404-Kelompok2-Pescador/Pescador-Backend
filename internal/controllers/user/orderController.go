package user

import (
	"Pescador-Backend/config"
	"Pescador-Backend/domain/entity"
	"net/http"

	"Pescador-Backend/internal/dto"

	"github.com/gofiber/fiber/v2"
)

/*
CreateOrder is a function to create shipping address
Get user cart
user just can create shipping address if user cart is not empty
Body request:
{
	"shipping_method": "JNE"
}

Response:
{
	"message": "Shipping address created",
	"data": {
		"id": "1",
		"shipping_method": "JNE",
		"shipping_price": 10000,
		"total_price": 10000,
		"created_at": "2021-08-01T00:00:00Z"
	}
}
*/
func CreateOrder(c *fiber.Ctx) error {
	user := c.Locals("user").(entity.UserToken)

	req := dto.OrderRequest{}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	var cart []entity.Cart

	err := config.DB.Where("user_id = ?", user.UserID).Find(&cart).Error
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Cart is empty",
		})
	}

	var shippingPrice float64

	if req.ShippingMethod == "JNE" {
		shippingPrice = 10000
	} else if req.ShippingMethod == "JNT" {
		shippingPrice = 15000
	} else if req.ShippingMethod == "Anteraja" {
		shippingPrice = 20000
	} else if req.ShippingMethod == "Same Day" {
		shippingPrice = 25000
	} else {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid shipping method",
		})
	}

	var totalPrice float64

	// get cart total price
	for _, cartItem := range cart {
		totalPrice += cartItem.TotalPrice + shippingPrice
	}

	// get store id from cart item
	// preloading product
	err = config.DB.Preload("Product").Where("user_id = ?", user.UserID).Find(&cart).Error

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to create order",
		})
	}

	// preloading store
	err = config.DB.Preload("Store").Where("id = ?", cart[0].Product.StoreID).Find(&cart[0].Product).Error

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to create order",
		})
	}

	order := entity.Order{
		UserID:         user.UserID,
		ShippingMethod: req.ShippingMethod,
		ShippingPrice:  shippingPrice,
		TotalPrice:     totalPrice,
		Status:         "Paid: Waiting for confirmation",
		StoreID:        cart[0].Product.StoreID,
	}

	// check user balance
	// if balance < TotalPrice then user can't create order
	var userBalance entity.UserBalance

	err = config.DB.Where("user_id = ?", user.UserID).Find(&userBalance).Error

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to create order Your balance is not enough",
		})
	}

	if userBalance.Balance < totalPrice {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to create order Your balance is not enough",
		})
	}

	// if sucess then create order
	err = config.DB.Create(&order).Error

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to create order",
		})
	}

	// Delete user cart
	err = config.DB.Where("user_id = ?", user.UserID).Delete(&cart).Error

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to create order",
		})
	}

	// update user balance
	userBalance.Balance -= totalPrice

	err = config.DB.Save(&userBalance).Error

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to create order",
		})
	}

	orderResponse := dto.OrderResponse{
		ID:             order.ID,
		ShippingMethod: order.ShippingMethod,
		ShippingPrice:  order.ShippingPrice,
		TotalPrice:     order.TotalPrice,
		Status:         order.Status,
		CreatedAt:      order.CreatedAt,
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Order created",
		"data":    orderResponse,
	})
}

/*
GetOrder is a function to get order by user id
Response:
{
	"message": "Your order",
	"data": [
		{
			"id": "1",
			"shipping_method": "JNE",
			"shipping_price": 10000,
			"total_price": 10000,
			"created_at": "2021-08-01T00:00:00Z"
		}
	]
}
*/
func GetOrder(c *fiber.Ctx) error {
	user := c.Locals("user").(entity.UserToken)

	var orders []entity.Order

	err := config.DB.Where("user_id = ?", user.UserID).Find(&orders).Error
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to get order",
		})
	}

	var orderResponses []dto.OrderResponse

	for _, v := range orders {
		orderResponse := dto.OrderResponse{
			ID:             v.ID,
			ShippingMethod: v.ShippingMethod,
			ShippingPrice:  v.ShippingPrice,
			TotalPrice:     v.TotalPrice,
			Status:         v.Status,
			CreatedAt:      v.CreatedAt,
		}

		orderResponses = append(orderResponses, orderResponse)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Your order",
		"data":    orderResponses,
	})
}
