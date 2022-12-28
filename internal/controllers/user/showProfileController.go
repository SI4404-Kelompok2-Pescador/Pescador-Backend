package user

import (
	"Pescador-Backend/config"
	"Pescador-Backend/domain/entity"
	"Pescador-Backend/internal/dto"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func ShowProfile(c *fiber.Ctx) error {
	user := c.Locals("user").(entity.UserToken)

	var userEntity entity.User

	err := config.DB.Where("id = ?", user.UserID).First(&userEntity).Error

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	profile := dto.Profile{
		Name:    userEntity.Name,
		Email:   userEntity.Email,
		Phone:   userEntity.Phone,
		Address: userEntity.Address,
		Picture: userEntity.Image,
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data":    profile,
	})
}
