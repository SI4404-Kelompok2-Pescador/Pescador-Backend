package routes

import (
	"Pescador-Backend/internal/controllers/admin"
	"Pescador-Backend/internal/controllers/store"
	"Pescador-Backend/internal/controllers/user"
	"Pescador-Backend/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	api := app.Group("/api")

	// =================== AUTH ===================
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
	// =================== AUTH ===================

	// =================== STORE ===================
	storeAPI := api.Group("/store").Use(middleware.AuthUser(middleware.Config{
		Unauthorized: func(c *fiber.Ctx) error {
			return c.Status(401).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		},
	}))
	storeAPI.Post("/create", store.RegisterStore)

	Store := api.Group("/login-store")
	Store.Post("", store.LoginStore)
	// =================== STORE ===================

	// =================== ADMIN ===================
	adminAPI := api.Group("/admin").Use(middleware.AuthAdmin(middleware.Config{
		Unauthorized: func(c *fiber.Ctx) error {
			return c.Status(401).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		},
	}))
	adminAPI.Get("/store", admin.ShowAllStore)
	// =================== ADMIN ===================

}
