package routes

import (
	"Pescador-Backend/internal/controllers/user"
	"Pescador-Backend/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	api := app.Group("/api")

	register := api.Group("/register")
	register.Post("", user.Register)

	login := api.Group("/login")
	login.Post("", user.Login)
	
	logout := api.Group("/logout").Use(middleware.AuthUser(middleware.Config{
		Unauthorized: func(c *fiber.Ctx) error {
			return c.Status(401).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		},
	}))
	logout.Post("", user.Logout)
}