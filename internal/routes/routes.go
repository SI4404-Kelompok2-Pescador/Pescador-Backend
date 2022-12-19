package routes

import (
	"Pescador-Backend/internal/controllers/user"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	api := app.Group("/api")

	register := api.Group("/register")
	register.Post("", user.Register)

	login := api.Group("/login")
	login.Post("", user.Login)
	
}