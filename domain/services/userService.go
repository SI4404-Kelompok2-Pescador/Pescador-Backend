package services

import (
	"github.com/gofiber/fiber/v2"
)

type UserService interface {
	Login(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
	Logout(c *fiber.Ctx) error
	ShowProfile(c *fiber.Ctx) error
	UpdateProfile(c *fiber.Ctx) error
	TopUpBalance(c *fiber.Ctx) error
	GetBalance(c *fiber.Ctx) error
	AddToCart(c *fiber.Ctx) error
	ViewCart(c *fiber.Ctx) error
	CreateOrder(c *fiber.Ctx) error
	GetOrder(c *fiber.Ctx) error
	AddWishlist(c *fiber.Ctx) error
	ShowWishlist(c *fiber.Ctx) error
}

