package user

import (
	"Pescador-Backend/config"
	"Pescador-Backend/domain/entity"
	"net/http"

	"Pescador-Backend/internal/dto"
	"github.com/gofiber/fiber/v2"
)

func TopUpBalance(c *fiber.Ctx) error {
	req := dto.UserBalanceRequest{}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	var userBalance entity.UserBalance
	
	user := c.Locals("user").(entity.UserToken)

	err := config.DB.Where("user_id = ?", user.UserID).First(&userBalance).Error

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Invalid Token",
		})
	}

	userBalance.Balance += req.Balance

	err = config.DB.Save(&userBalance).Error

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Could not top up balance",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Balance top up success",
		"data": dto.UserBalanceResponse{
			Balance: userBalance.Balance,
		},
	})

}

func GetBalance(c *fiber.Ctx) error {
	var userBalance entity.UserBalance

	user := c.Locals("user").(entity.UserToken)

	err := config.DB.Where("user_id = ?", user.UserID).First(&userBalance).Error

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Invalid Token",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Balance found",
		"data": dto.UserBalanceResponse{
			Balance: userBalance.Balance,
		},
	})
}