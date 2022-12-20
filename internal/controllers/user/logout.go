package user

import (
	"net/http"
	// "time"
	// "fmt"

	"Pescador-Backend/internal/database"
	"Pescador-Backend/internal/models/auth"
	"Pescador-Backend/internal/models/user"

	"github.com/gofiber/fiber/v2"
)

func Logout(c *fiber.Ctx) error {
	// Get user ID from JWT token
	user_id := c.Locals("user").(user.User)

	token := auth.UserToken{}

	err := database.DB.Where("user_id = ?", user_id.ID).Delete(&token).Error

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not logout",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Logged out",
	})

}
