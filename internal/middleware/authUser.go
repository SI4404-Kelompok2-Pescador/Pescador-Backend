package middleware

import (
	"log"

	"Pescador-Backend/internal/database"
	"Pescador-Backend/internal/models/auth"
	"github.com/gofiber/fiber/v2"
)

func AuthUser(c Config) fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		header := ctx.GetReqHeaders()

		if _, ok := header["Authorization"]; !ok {
			return c.Unauthorized(ctx)
		}

		userToken := auth.UserToken{}
		err := database.DB.Where("token = ?", header["Authorization"]).First(&userToken).Error
		if err != nil {
			return c.Unauthorized(ctx)
		}

		if userToken.Type != "buyer" {
			return c.Unauthorized(ctx)
		}

		ctx.Locals("user", userToken.User)
		log.Println("User Authenticated")
		return ctx.Next()
	}

}