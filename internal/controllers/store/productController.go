package store

import (
	"net/http"

	"Pescador-Backend/config"
	"Pescador-Backend/domain/entity"
	"Pescador-Backend/internal/dto"

	"github.com/gofiber/fiber/v2"
)

func (s *StoreImplementation) CreateProduct(c *fiber.Ctx) error {
	req := dto.ProductRequest{}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	store := c.Locals("store").(entity.StoreToken)

	// check if category name exists
	var category entity.Category

	err := config.DB.Where("name = ?", req.CategoryName).First(&category).Error

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Category name does not exist",
		})
	}

	// check if category name is exist in database
	categoryName := entity.Category{}

	err = config.DB.Where("name = ?", req.CategoryName).First(&categoryName).Error

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Category name does not exist",
		})
	}

	newProduct := entity.Product{
		Name:         req.Name,
		Price:        req.Price,
		Stock:        req.Stock,
		Description:  req.Description,
		Image:        req.Image,
		CategoryName: req.CategoryName,
		StoreID:      store.StoreID,
	}

	err = config.DB.Create(&newProduct).Error

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	// preload store name
	err = config.DB.Preload("Store").Where("id = ?", newProduct.ID).First(&newProduct).Error

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	// preload category name
	err = config.DB.Preload("Category").Where("id = ?", newProduct.ID).First(&newProduct).Error

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	response := dto.ProductResponse{
		ID:          newProduct.ID,
		Name:        newProduct.Name,
		Price:       newProduct.Price,
		Stock:       newProduct.Stock,
		Description: newProduct.Description,
		Image:       newProduct.Image,
		Category:    newProduct.Category.Name,
		StoreName:   newProduct.Store.Name,
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "success",
		"data":    response,
	})

}

func (s *StoreImplementation) DetailsProduct(c *fiber.Ctx) error {
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
