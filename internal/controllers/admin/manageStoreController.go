package admin

import (
	"net/http"

	"Pescador-Backend/internal/database"
	"Pescador-Backend/internal/dto"
	"Pescador-Backend/internal/models"

	"github.com/gofiber/fiber/v2"
)

func ShowAllStore(c *fiber.Ctx) error {
	store := []models.Store{}

	// show all store in database with owner and don't show owner password
	err := database.DB.Find(&store).Error

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not get store",
		})
	}

	storeResponse := []dto.StoreResponse{}

	for _, store := range store {
		storeResponse = append(storeResponse, dto.StoreResponse{
			ID:      store.ID,
			Name:    store.Name,
			Email:   store.Email,
			Phone:   store.Phone,
			Address: store.Address,
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Store found",
		"data":    storeResponse,
	})

}
