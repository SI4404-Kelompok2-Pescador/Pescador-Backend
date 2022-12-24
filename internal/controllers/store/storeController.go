package store

import (
	"net/http"
	"time"

	"Pescador-Backend/internal/database"
	"Pescador-Backend/internal/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
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

	pass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	var storeType models.Type

	err = database.DB.Where("name = ?", "store").First(&storeType).Error

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	err = database.DB.Create(&storeType).Error

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	stores := models.Store{}

	err = database.DB.Where("user_id = ?", user.UserID).First(&stores).Error

	if err == nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Store already registered",
		})
	}

	newStore := models.Store{
		Name:     req.Name,
		Email:    req.Email,
		Phone:    req.Phone,
		Address:  req.Address,
		Password: string(pass),
		OwnerID:  user.UserID,
		Type:     storeType.Name,
	}

	// check if store already exists
	var existingStore models.Store
	err = database.DB.Where("email = ?", newStore.Email).First(&existingStore).Error
	if err == nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Store already exists",
		})
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

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Store registered",
		"data":    response,
	})

}

func LoginStore(c *fiber.Ctx) error {
	req := models.StoreLoginRequest{}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	store := models.Store{}

	err := database.DB.Where("email = ?", req.Email).First(&store).Error

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Store not found",
		})
	}

	err = bcrypt.CompareHashAndPassword([]byte(store.Password), []byte(req.Password))

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid password",
		})
	}

	// Generate JWT token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    store.Name,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
	})

	token, err := claims.SignedString([]byte("secret"))

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error signing token",
		})
	}

	storeToken := models.StoreToken{
		StoreID: store.ID,
		Type:    store.Type,
		Token:   token,
	}

	err = database.DB.Create(&storeToken).Error

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	response := models.StoreLoginResponse{
		ID:      store.ID,
		Name:    store.Name,
		Email:   store.Email,
		Phone:   store.Phone,
		Address: store.Address,
		OwnerID: store.OwnerID,
		Type:    store.Type,
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Store logged in",
		"token":   token,
		"data":    response,
	})

}
