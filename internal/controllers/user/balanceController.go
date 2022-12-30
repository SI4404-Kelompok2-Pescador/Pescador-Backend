package user

import (
	"Pescador-Backend/config"
	"Pescador-Backend/domain/entity"
	"net/http"

	"Pescador-Backend/internal/dto"

	"github.com/gofiber/fiber/v2"
)

func TopUpBalance(c *fiber.Ctx) error {
	user := c.Locals("user").(entity.UserToken)

	req := dto.UserBalanceRequest{}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	var userBalance entity.UserBalance

	// create new balance if user first time top up
	// if user first time top up, create new balance
	if userBalance.ID == "" {
		userBalance = entity.UserBalance{
			UserID:  user.UserID,
			Balance: req.Balance,
		}
		err := config.DB.Create(&userBalance).Error
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to top up balance",
			})
		}
	} else {
		userBalance.Balance += req.Balance
		err := config.DB.Save(&userBalance).Error
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to top up balance",
			})
		}

	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Balance updated",
		"data": dto.UserBalanceResponse{
			Balance: userBalance.Balance,
		},
	})

}

func GetBalance(c *fiber.Ctx) error {
	var userBalance entity.UserBalance

	user := c.Locals("user").(entity.UserToken)

	err := config.DB.Where("user_id = ?", user.UserID).First(&userBalance).Error

	// if user id is not found, return 0 balance
	if err != nil {
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"message": "Balance Not Found Please Top Up",
			"data": dto.UserBalanceResponse{
				Balance: 0,
			},
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Your Balance",
		"data": dto.UserBalanceResponse{
			Balance: userBalance.Balance,
		},
	})
}
