package middleware

import (
	"github.com/gofiber/fiber/v2"
)

type Config struct {
	Filter       func(*fiber.Ctx) error
	Unauthorized fiber.Handler
}