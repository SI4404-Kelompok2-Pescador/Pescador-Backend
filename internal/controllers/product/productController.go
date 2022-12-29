package product

import (
	"net/http"

	"Pescador-Backend/config"
	"Pescador-Backend/domain/entity"
	"Pescador-Backend/internal/dto"

	"github.com/gofiber/fiber/v2"
)

func CreateProduct(c *fiber.Ctx) error {
	req := dto.ProductRequest{}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	// product := entity.Product{}

	store := c.Locals("store").(entity.StoreToken)

	// err := config.DB.Where("store_id = ?", store.StoreID).First(&product).Error

	// if err != nil {
	// 	return c.Status(http.StatusInternalServerError).JSON(err.Error())
	// }

	newProduct := entity.Product{
		Name:        req.Name,
		Price:       req.Price,
		Stock:       req.Stock,
		Description: req.Description,
		Image:       req.Image,
		StoreID:     store.StoreID,
	}

	err := config.DB.Create(&newProduct).Error

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	// preload store name
	err = config.DB.Preload("Store").Where("id = ?", newProduct.ID).First(&newProduct).Error

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
		StoreName:   newProduct.Store.Name,
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data":    response,
	})

}
