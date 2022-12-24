package main

import (
	"os"

	"Pescador-Backend/config"
	"Pescador-Backend/internal/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load(".env")

	if err != nil {
		panic("Error loading .env file")
	}

	// Connect to config
	config.Connect(&config.Config{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
		SSLMODE:  os.Getenv("DB_SSLMODE"),
	})

	// Create Fiber app
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Token, Type",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowCredentials: true,
	}))

	routes.Setup(app)
	err = app.Listen(":" + os.Getenv("PORT"))

	if err != nil {
		panic(err)
	}
}
