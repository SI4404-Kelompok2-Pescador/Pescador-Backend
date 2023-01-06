package services

import (
	"github.com/gofiber/fiber/v2"
)

type StoreService interface {
	RegisterStore(c *fiber.Ctx) error
	LoginStore(c *fiber.Ctx) error
	GetStoreProducts(c *fiber.Ctx) error
	GetOrder(c *fiber.Ctx) error
	UpdateOrder(c *fiber.Ctx) error
	CreateProduct(c *fiber.Ctx) error
	DetailsProduct(c *fiber.Ctx) error
}
