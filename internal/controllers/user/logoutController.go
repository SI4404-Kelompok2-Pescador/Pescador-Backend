package user

import (
	"Pescador-Backend/config"
	"Pescador-Backend/domain/entity"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

/*
Logout is a function to logout user
Response:
{
	"status": "success",
	"message": "Logged out"
}
*/
func Logout(c *fiber.Ctx) error {
	// Get user ID from JWT token
	user := c.Locals("user").(entity.UserToken)

	token := entity.UserToken{}

	err := config.DB.Where("user_id = ?", user.UserID).Delete(&token).Error

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not logout",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Logged out",
	})

}
