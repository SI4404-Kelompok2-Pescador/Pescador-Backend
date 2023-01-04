package user

import (
	"Pescador-Backend/config"
	"Pescador-Backend/domain/entity"
	"net/http"

	"Pescador-Backend/internal/dto"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func (u * UserImplementation) UpdateProfile(c *fiber.Ctx) error {
	req := dto.UpdateProfileRequest{}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	user := c.Locals("user").(entity.UserToken)

	var userEntity entity.User

	err := config.DB.Where("id = ?", user.UserID).First(&userEntity).Error

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	// update user
	userEntity.Name = req.Name
	userEntity.Email = req.Email
	userEntity.Phone = req.Phone
	userEntity.Address = req.Address
	userEntity.Image = req.Image

	// password cannot be same as before
	// check the request password with the password in database
	err = bcrypt.CompareHashAndPassword([]byte(userEntity.Password), []byte(req.Password))

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Password cannot be same as before",
		})
	}

	// if password is different, then update the password
	pass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	userEntity.Password = string(pass)

	err = config.DB.Save(&userEntity).Error

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Profile updated",
	})
}
