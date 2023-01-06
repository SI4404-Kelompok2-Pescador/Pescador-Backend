package services

import (
	"github.com/gofiber/fiber/v2"
)

type GlobalService interface {
	DetailsProduct(c *fiber.Ctx) error
	GetAllCategories(c *fiber.Ctx) error
	ShowAllProduct(c *fiber.Ctx) error
}
