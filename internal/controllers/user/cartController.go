package user

import (
	"Pescador-Backend/config"
	"Pescador-Backend/domain/entity"
	"net/http"

	"Pescador-Backend/internal/dto"

	"github.com/gofiber/fiber/v2"
)

/*
AddToCart is a function to add product to cart
Body request:
{
	"product_id": "1",
	"quantity": 1
}
Response:
{
	"message": "Product added to cart",
	"data": {
		"id": "1",
		"product_id": "1",
		"quantity": 1,
		"price": 10000
	}
}
*/
func AddToCart(c *fiber.Ctx) error {
	user := c.Locals("user").(entity.UserToken)

	req := dto.CartRequest{}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	var product entity.Product

	err := config.DB.Where("id = ?", req.ProductID).First(&product).Error
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Product not found",
		})
	}

	var cart entity.Cart

	// if product already in cart, update quantity
	// get productID in user cart
	err = config.DB.Where("user_id = ?", user.UserID).Where("product_id = ?", req.ProductID).First(&cart).Error

	if err != nil {
		// if product not in cart, create new cart
		cart = entity.Cart{
			UserID:     user.UserID,
			ProductID:  req.ProductID,
			Quantity:   req.Quantity,
			TotalPrice: product.Price * float64(req.Quantity),
		}

		err = config.DB.Create(&cart).Error

		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to add product to cart",
			})
		}
	} else {
		// if product already in cart, update quantity
		cart.Quantity = cart.Quantity + req.Quantity
		cart.TotalPrice = cart.TotalPrice + (product.Price * float64(req.Quantity))

		err = config.DB.Save(&cart).Error

		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to add product to cart",
			})
		}
	}

	cartResponse := dto.CartResponse{
		ID:           cart.ID,
		ProductName:  product.Name,
		Quantity:     cart.Quantity,
		ProductPrice: product.Price,
		TotalPrice:   cart.TotalPrice,
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Product added to cart",
		"status":  "success",
		"data":    cartResponse,
	})

}

/*
ViewCart is a function to view cart
Response:
{
	"message": "Cart list",
	"data": [
		{
			"id": "1",
			"product_name": "Product 1",
			"quantity": 1,
			"price": 10000
		},
		{
			"id": "2",
			"product_name": "Product 2",
			"quantity": 1,
			"price": 20000
		}
	]
}
*/
func ViewCart(c *fiber.Ctx) error {
	user := c.Locals("user").(entity.UserToken)

	var carts []entity.Cart

	err := config.DB.Where("user_id = ?", user.UserID).Find(&carts).Error

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get cart list",
		})
	}

	var cartList []dto.CartListResponse

	for _, cart := range carts {
		var product entity.Product

		err = config.DB.Where("id = ?", cart.ProductID).First(&product).Error

		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to get cart list",
			})
		}

		cartList = append(cartList, dto.CartListResponse{
			ID:           cart.ID,
			ProductName:  product.Name,
			Quantity:     cart.Quantity,
			ProductPrice: product.Price,
			TotalPrice:   cart.TotalPrice,
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Cart list",
		"status":  "success",
		"data":    cartList,
	})
}
