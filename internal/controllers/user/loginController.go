package user

import (
	"Pescador-Backend/config"
	"Pescador-Backend/domain/entity"
	"net/http"
	"os"
	"time"

	"Pescador-Backend/internal/dto"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

/*
GetUserToken is a function to get user token from .env file
*/
func GetUserToken() string {
	return os.Getenv("USER_TOKEN_SECRET")
}

/*
Login is a function to login user
Body request:
{
	"email": "
	"password": "
}
Response:
{
	"message": "Login success",
	"data": {
		"name": "John Doe",
		"email": "
		"phone": "
		"address": "
	},
	"token": "
}
*/
func (u *UserImplementation) Login(c *fiber.Ctx) error {
	req := dto.UserLoginRequest{}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	userLogin := entity.User{}

	err := config.DB.Where("email = ?", req.Email).First(&userLogin).Error
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
	claims := dto.CustomClaims{
		UserID:   userLogin.ID,
		UserName: userLogin.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(GetUserToken()))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not login",
		})
	}

	userToken := entity.UserToken{
		UserID: userLogin.ID,
		Type:   userLogin.Type,
		Token:  signedToken,
	}

	// Save JWT token to database
	// if user has login before, update token
	// if user has not login before, create new token
	err = config.DB.Where("user_id = ?", userLogin.ID).First(&userToken).Error
	if err != nil {
		config.DB.Create(&userToken)
	} else {
		config.DB.Model(&userToken).Where("user_id = ?", userLogin.ID).Update("token", signedToken)
	}

	// delete token after reach expire time
	go func() {
		if token.Valid {
			time.Sleep(time.Hour * 24 * 7)
			config.DB.Delete(&userToken)
		}
	}()

	response := dto.LoginResponse{
		ID:      userLogin.ID,
		Name:    userLogin.Name,
		Email:   userLogin.Email,
		Phone:   userLogin.Phone,
		Address: userLogin.Address,
		Type:    userLogin.Type,
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Login Successful",
		"token":   signedToken,
		"user":    response,
	})

}
