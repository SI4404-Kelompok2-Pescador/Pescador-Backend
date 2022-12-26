package routes

import (
	"Pescador-Backend/internal/controllers/admin"
	"Pescador-Backend/internal/controllers/store"
	"Pescador-Backend/internal/controllers/product"
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

	users := api.Group("/user").Use(middleware.AuthUser(middleware.Config{
		Unauthorized: func(c *fiber.Ctx) error {
			return c.Status(401).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		},
	}))
	users.Post("/logout", user.Logout)
	users.Put("/profile", user.UpdateProfile)
	users.Get("/profile", user.ShowProfile)
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

	// =================== Product ===================
	productAPI := api.Group("/product").Use(middleware.AuthStore(middleware.Config{
		Unauthorized: func(c *fiber.Ctx) error {
			return c.Status(401).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		},
	}))
	productAPI.Post("/create", product.CreateProduct)

	// =================== ADMIN ===================
	adminAPI := api.Group("/admin").Use(middleware.AuthAdmin(middleware.Config{
		Unauthorized: func(c *fiber.Ctx) error {
			return c.Status(401).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		},
	}))
	adminAPI.Get("/store", admin.ShowAllStore)
	adminAPI.Get("/store/:id", admin.GetStoreByID)
	// =================== ADMIN ===================

}
