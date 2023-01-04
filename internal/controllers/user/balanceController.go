package user

import (
	"Pescador-Backend/config"
	"Pescador-Backend/domain/entity"
	"net/http"

	"Pescador-Backend/internal/dto"

	"github.com/gofiber/fiber/v2"
)

/*
TopUpBalance is a function to top up user balance
Body request:
{
	"balance": 10000
}
Response:
{
	"message": "Balance updated",
	"data": {
		"balance": 10000
	}
}
*/
func TopUpBalance(c *fiber.Ctx) error {
	user := c.Locals("user").(entity.UserToken)

	req := dto.UserBalanceRequest{}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	var userBalance entity.UserBalance

	err := config.DB.Where("user_id = ?", user.UserID).First(&userBalance).Error

	if err != nil {
		// if user id is not found, create new user balance
		userBalance = entity.UserBalance{
			UserID:  user.UserID,
			Balance: req.Balance,
		}

		err = config.DB.Create(&userBalance).Error

		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"message": "Failed to create balance",
			})
		}
	} else {
		// if user id is found, update balance
		userBalance.Balance += req.Balance

		err = config.DB.Save(&userBalance).Error

		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"message": "Failed to update balance",
			})
		}
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Balance updated",
		"status":  "success",
		"data": dto.UserBalanceResponse{
			Balance: userBalance.Balance,
		},
	})

}

/*
GetBalance is a function to get user balance
Response:
{
	"message": "Your Balance",
	"data": {
		"balance": 10000
	}
}
*/
func (u * UserImplementation) GetBalance(c *fiber.Ctx) error {
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
		"status":  "success",
		"data": dto.UserBalanceResponse{
			Balance: userBalance.Balance,
		},
	})
}
