package admin

import (
	"Pescador-Backend/config"
	"Pescador-Backend/domain/entity"
	"net/http"

	"Pescador-Backend/internal/dto"

	"github.com/gofiber/fiber/v2"
)

func (a *AdminImplementation) CreateCategory(c *fiber.Ctx) error {
	req := dto.CategoryRequest{}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	newCategory := entity.Category{
		Name: req.Name,
	}

	err := config.DB.Create(&newCategory).Error

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	response := dto.CategoryResponse{
		ID:   newCategory.ID,
		Name: newCategory.Name,
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data":    response,
	})
}

func (a *AdminImplementation) GetAllCategories(c *fiber.Ctx) error {
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


