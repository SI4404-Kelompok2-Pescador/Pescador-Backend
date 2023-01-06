package global

import (
	"net/http"

	"Pescador-Backend/config"
	"Pescador-Backend/domain/entity"
	"Pescador-Backend/internal/dto"

	"github.com/gofiber/fiber/v2"
)

func (g *GlobalImplementation) DetailsProduct(c *fiber.Ctx) error {
	id := c.Params("id")

	product := entity.Product{}

	err := config.DB.Preload("Store").Where("id = ?", id).First(&product).Error

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	response := dto.ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Price:       product.Price,
		Stock:       product.Stock,
		Description: product.Description,
		Image:       product.Image,
		StoreName:   product.Store.Name,
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data":    response,
	})
}

func (g *GlobalImplementation) GetAllCategories(c *fiber.Ctx) error {
	var categories []entity.Category

	err := config.DB.Find(&categories).Error

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	var response []dto.CategoryResponse

	for _, category := range categories {
		response = append(response, dto.CategoryResponse{
			ID:   category.ID,
			Name: category.Name,
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data":    response,
	})
}

func (g *GlobalImplementation) ShowAllProduct(c *fiber.Ctx) error {
	var product []entity.Product

	// show all product in config with store and don't show store password
	err := config.DB.Find(&product).Error

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not get product",
		})
	}

	// preload store name and category name
	err = config.DB.Preload("Store").Preload("Category").Find(&product).Error

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
			Image:       product.Image,
			Category:    product.Category.Name,
			StoreName:   product.Store.Name,
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Product found",
		"data":    productResponse,
	})

}
