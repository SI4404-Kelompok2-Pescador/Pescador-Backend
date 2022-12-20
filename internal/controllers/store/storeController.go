package store

import (
	"Pescador-Backend/internal/database"
	"Pescador-Backend/internal/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func RegisterStore(c *fiber.Ctx) error {
	// Get userID from JWT token
	user := c.Locals("user").(models.UserToken)

	req := models.StoreRequest{}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	stores := models.Store{}

	err := database.DB.Where("user_id = ?", user.UserID).First(&stores).Error

	if err == nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Store already registered",
		})
	}

	newStore := models.Store{
		Name:    req.Name,
		Email:   req.Email,
		Phone:   req.Phone,
		Address: req.Address,
		OwnerID: user.UserID,
	}

	err = database.DB.Create(&newStore).Error

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	response := models.StoreResponse{
		ID:      newStore.ID,
		Name:    newStore.Name,
		Email:   newStore.Email,
		Phone:   newStore.Phone,
		Address: newStore.Address,
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Store registered",
		"data":    response,
	})

}
