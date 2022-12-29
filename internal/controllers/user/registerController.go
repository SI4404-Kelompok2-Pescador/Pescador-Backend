package user

import (
	"Pescador-Backend/config"
	"Pescador-Backend/domain/entity"
	"net/http"

	"Pescador-Backend/internal/dto"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	req := dto.UserRequest{}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	var buyerType entity.Type

	err = config.DB.Where("name = ?", "buyer").First(&buyerType).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	newUser := entity.User{
		Name:     req.Name,
		Email:    req.Email,
		Phone:    req.Phone,
		Address:  req.Address,
		Image:    req.Image,
		Password: string(hashedPassword),
		Type:     buyerType.Name,
	}

	// check if user already exists
	var existingUser entity.User
	err = config.DB.Where("email = ?", newUser.Email).First(&existingUser).Error
	if err == nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "User already exists",
		})
	}

	err = config.DB.Create(&newUser).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	newUserType := entity.UserType{
		UserID: newUser.ID,
		TypeID: buyerType.ID,
	}

	err = config.DB.Create(&newUserType).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "User created successfully",
		"data":    newUser,
	})

}
