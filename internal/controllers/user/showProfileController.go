package user

import (
	"Pescador-Backend/config"
	"Pescador-Backend/domain/entity"
	"Pescador-Backend/internal/dto"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

/*
ShowProfile is a function to show user profile
Response:
{
	"message": "success",
	"data": {
		"name": "John Doe",
		"email": "",
		"phone": "",
		"address": "",
		"image": ""
	}
}
*/
func (u *UserImplementation) ShowProfile(c *fiber.Ctx) error {
	user := c.Locals("user").(entity.UserToken)

	var userEntity entity.User

	err := config.DB.Where("id = ?", user.UserID).First(&userEntity).Error

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	profile := dto.Profile{
		ID:      userEntity.ID,
		Name:    userEntity.Name,
		Email:   userEntity.Email,
		Phone:   userEntity.Phone,
		Address: userEntity.Address,
		Image:   userEntity.Image,
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data":    profile,
	})
}
