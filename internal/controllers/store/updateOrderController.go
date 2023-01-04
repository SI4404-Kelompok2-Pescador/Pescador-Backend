package store

import (
	"Pescador-Backend/config"
	"Pescador-Backend/domain/entity"
	"net/http"

	"Pescador-Backend/internal/dto"

	"github.com/gofiber/fiber/v2"
)

func GetOrder(c *fiber.Ctx) error {
	// Get order that belongs to store
	store := c.Locals("store").(entity.StoreToken)

	var orders []entity.Order

	err := config.DB.Where("store_id = ?", store.StoreID).Find(&orders).Error

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Order not found",
		})
	}

	/*
		StoreOrderResponse struct {
			ID             string    `json:"id"`
			ShippingMethod string    `json:"shipping_method"`
			ShippingPrice  float64   `json:"shipping_price"`
			TotalPrice     float64   `json:"total_price"`
			UserName       string    `json:"user_name"`
			UserAddress    string    `json:"user_address"`
			UserPhone      string    `json:"user_phone"`
			UserEmail      string    `json:"user_email"`
			Status         string    `json:"status"`
			CreatedAt      time.Time `json:"created_at"`
		}
	*/

	var storeOrders []dto.StoreOrderResponse

	for _, order := range orders {
		var user entity.User

		err := config.DB.Where("id = ?", order.UserID).First(&user).Error

		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"message": "User not found",
			})
		}

		storeOrder := dto.StoreOrderResponse{
			ID:             order.ID,
			ShippingMethod: order.ShippingMethod,
			ShippingPrice:  order.ShippingPrice,
			TotalPrice:     order.TotalPrice,
			UserName:       user.Name,
			UserAddress:    user.Address,
			UserPhone:      user.Phone,
			UserEmail:      user.Email,
			Status:         order.Status,
			CreatedAt:      order.CreatedAt,
		}

		storeOrders = append(storeOrders, storeOrder)

	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Order found",
		"status":  "success",
		"data":    storeOrders,
	})
}

func UpdateOrder(c *fiber.Ctx) error {
	orderID := c.Query("id")

	req := dto.UpdateOrderRequest{}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	var order entity.Order

	err := config.DB.Where("id = ?", orderID).First(&order).Error

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Order not found",
		})
	}

	order.Status = req.Status

	err = config.DB.Save(&order).Error

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	var user entity.User

	err = config.DB.Where("id = ?", order.UserID).First(&user).Error

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	response := dto.StoreOrderResponse{
		ID:             order.ID,
		ShippingMethod: order.ShippingMethod,
		ShippingPrice:  order.ShippingPrice,
		TotalPrice:     order.TotalPrice,
		UserName:       user.Name,
		UserAddress:    user.Address,
		UserPhone:      user.Phone,
		UserEmail:      user.Email,
		Status:         order.Status,
		CreatedAt:      order.CreatedAt,
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Order updated",
		"status":  "success",
		"data":    response,
	})
}
