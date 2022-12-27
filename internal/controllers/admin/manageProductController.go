package admin

import (
	"Pescador-Backend/config"
	"Pescador-Backend/domain/entity"
	"net/http"

	"Pescador-Backend/internal/dto"

	"github.com/gofiber/fiber/v2"
)

func ShowAllProduct(c *fiber.Ctx) error {
	var product []entity.Product

	// show all product in config with store and don't show store password
	err := config.DB.Find(&product).Error

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not get product",
		})
	}

	// preload store name
	err = config.DB.Preload("Store").Find(&product).Error

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not get product",
		})
	}

	var productResponse []dto.ProductResponse

	for _, product := range product {
		productResponse = append(productResponse, dto.ProductResponse{
			ID:          product.ID,
			Name:        product.Name,
			Price:       product.Price,
			Stock:       product.Stock,
			Description: product.Description,
			Picture:     product.Picture,
			StoreName:   product.Store.Name,
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Product found",
		"data":    productResponse,
	})

}
