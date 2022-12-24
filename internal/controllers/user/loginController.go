package user

import (
	"net/http"
	"os"
	"time"

	"Pescador-Backend/internal/database"
	"Pescador-Backend/internal/dto"
	"Pescador-Backend/internal/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func GetUserToken() string {
	return os.Getenv("USER_TOKEN")
}

func Login(c *fiber.Ctx) error {
	req := dto.UserLoginRequest{}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	userLogin := models.User{}

	err := database.DB.Where("email = ?", req.Email).First(&userLogin).Error
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "User Not Found",
		})
	}

	err = bcrypt.CompareHashAndPassword([]byte(userLogin.Password), []byte(req.Password))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Credentials",
		})
	}

	// Create JWT token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    userLogin.ID,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
	})

	token, err := claims.SignedString([]byte(GetUserToken()))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not login",
		})
	}

	userToken := models.UserToken{
		UserID: userLogin.ID,
		Type:   userLogin.Type,
		Token:  token,
	}

	err = database.DB.Create(&userToken).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	response := dto.LoginResponse{
		Name:    userLogin.Name,
		Email:   userLogin.Email,
		Phone:   userLogin.Phone,
		Address: userLogin.Address,
		Type:    userLogin.Type,
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Login Successful",
		"token":   token,
		"user":    response,
	})

}
