package routes

import (
	"Pescador-Backend/internal/controllers/admin"
	"Pescador-Backend/internal/controllers/product"
	"Pescador-Backend/internal/controllers/store"
	"Pescador-Backend/internal/controllers/user"
	"Pescador-Backend/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	api := app.Group("/api")
	userImplementation := user.UserImplementation{}

	// =================== AUTH ===================
	register := api.Group("/register")
	register.Post("", userImplementation.Register)

	login := api.Group("/login")
	login.Post("", userImplementation.Login)

	users := api.Group("/user").Use(middleware.AuthUser(middleware.Config{
		Unauthorized: func(c *fiber.Ctx) error {
			return c.Status(401).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		},
	}))
	users.Post("/logout", userImplementation.Logout)
	users.Put("/profile", userImplementation.UpdateProfile)
	users.Get("/profile", userImplementation.ShowProfile)
	// =================== AUTH ===================

	// ==================== Global ====================
	products := api.Group("/products")
	products.Get("/detail", product.DetailsProduct)
	products.Get("", admin.ShowAllProduct)
	categories := api.Group("/categories")
	categories.Get("", admin.GetAllCategories)
	// ==================== Global ====================

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
	// =================== Product ===================
	productAPI := api.Group("/product").Use(middleware.AuthStore(middleware.Config{
		Unauthorized: func(c *fiber.Ctx) error {
			return c.Status(401).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		},
	}))
	productAPI.Post("/create", product.CreateProduct)
	productAPI.Get("/shows", store.GetStoreProducts)

	// =================== Product ===================

	// =================== ORDER ===================
	orderAPI := api.Group("/order").Use(middleware.AuthStore(middleware.Config{
		Unauthorized: func(c *fiber.Ctx) error {
			return c.Status(401).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		},
	}))
	orderAPI.Get("", store.GetOrder)
	orderAPI.Put("/update", store.UpdateOrder)

	// =================== STORE ===================

	// =================== USER =============================
	userAPI := api.Group("/user").Use(middleware.AuthUser(middleware.Config{
		Unauthorized: func(c *fiber.Ctx) error {
			return c.Status(401).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		},
	}))

	// =================== BALANCE ===================
	balance := userAPI.Group("/balance")
	balance.Post("", user.TopUpBalance)
	balance.Get("", userImplementation.GetBalance)
	// =================== BALANCE ===================

	// =================== CART ===================
	cart := userAPI.Group("/cart")
	cart.Post("/add", user.AddToCart)
	cart.Get("/show", userImplementation.ViewCart)
	// =================== CART ===================

	// =================== ORDER ===================
	order := userAPI.Group("/order")
	order.Post("/create", userImplementation.CreateOrder)
	order.Get("", userImplementation.GetOrder)

	// =================== USER =============================

	// =================== ADMIN ===================
	adminAPI := api.Group("/admin").Use(middleware.AuthAdmin(middleware.Config{
		Unauthorized: func(c *fiber.Ctx) error {
			return c.Status(401).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		},
	}))
	adminAPI.Get("/stores", admin.ShowAllStore)
	// get store by id
	adminAPI.Get("/store", admin.GetStoreByID)
	// Create Category
	category := adminAPI.Group("/category")
	category.Post("/create", admin.CreateCategory)
	// =================== ADMIN ===================

}
