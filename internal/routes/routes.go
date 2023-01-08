package routes

import (
	"Pescador-Backend/internal/controllers/admin"
	"Pescador-Backend/internal/controllers/global"
	"Pescador-Backend/internal/controllers/store"
	"Pescador-Backend/internal/controllers/user"
	"Pescador-Backend/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	userImplementation := user.UserImplementation{}
	adminImplementation := admin.AdminImplementation{}
	storeImplementation := store.StoreImplementation{}
	globalImplementation := global.GlobalImplementation{}

	api := app.Group("/api")

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
	products.Get("/detail", globalImplementation.DetailsProduct)
	products.Get("", globalImplementation.ShowAllProduct)
	categories := api.Group("/categories")
	categories.Get("", globalImplementation.GetAllCategories)
	// ==================== Global ====================

	// =================== STORE ===================
	storeAPI := api.Group("/store").Use(middleware.AuthUser(middleware.Config{
		Unauthorized: func(c *fiber.Ctx) error {
			return c.Status(401).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		},
	}))
	storeAPI.Post("/create", storeImplementation.RegisterStore)

	Store := api.Group("/login-store")
	Store.Post("", storeImplementation.LoginStore)
	// =================== Product ===================
	productAPI := api.Group("/product").Use(middleware.AuthStore(middleware.Config{
		Unauthorized: func(c *fiber.Ctx) error {
			return c.Status(401).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		},
	}))
	productAPI.Post("/create", storeImplementation.CreateProduct)
	productAPI.Get("/shows", storeImplementation.GetStoreProducts)

	// =================== Product ===================

	// =================== ORDER ===================
	orderAPI := api.Group("/order").Use(middleware.AuthStore(middleware.Config{
		Unauthorized: func(c *fiber.Ctx) error {
			return c.Status(401).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		},
	}))
	orderAPI.Get("", storeImplementation.GetOrder)
	orderAPI.Put("/update", storeImplementation.UpdateOrder)

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
	// =================== ORDER ===================

	// =================== Wishlist ===================
	wishlist := userAPI.Group("/wishlist")
	wishlist.Post("/add", userImplementation.AddWishlist)
	wishlist.Get("/show", userImplementation.ShowWishlist)
	// =================== Wishlist ===================


	// =================== USER =============================

	// =================== ADMIN ===================
	adminAPI := api.Group("/admin").Use(middleware.AuthAdmin(middleware.Config{
		Unauthorized: func(c *fiber.Ctx) error {
			return c.Status(401).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		},
	}))
	adminAPI.Get("/stores", adminImplementation.ShowAllStore)
	// get store by id
	adminAPI.Get("/store", adminImplementation.GetStoreByID)
	// Create Category
	category := adminAPI.Group("/category")
	category.Post("/create", adminImplementation.CreateCategory)
	// =================== ADMIN ===================

}
