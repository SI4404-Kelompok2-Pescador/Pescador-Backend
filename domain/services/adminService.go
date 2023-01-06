package services

import (
	"github.com/gofiber/fiber/v2"
)

type AdminService interface {
	CreateCategory(c *fiber.Ctx) error
	GetAllCategories(c *fiber.Ctx) error
	ShowAllStores(c *fiber.Ctx) error
	GetStoreByID(c *fiber.Ctx) error
}