package middleware

import (
	"Pescador-Backend/config"
	"Pescador-Backend/domain/entity"
	"log"

	"github.com/gofiber/fiber/v2"
)

func AuthAdmin(c Config) fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		header := ctx.GetReqHeaders()

		if _, ok := header["Authorization"]; !ok {
			return c.Unauthorized(ctx)
		}

		userToken := entity.UserToken{}
		err := config.DB.Where("token = ?", header["Authorization"]).First(&userToken).Error
		if err != nil {
			return c.Unauthorized(ctx)
		}

		if userToken.Type != "admin" {
			return c.Unauthorized(ctx)
		}

		ctx.Locals("user", userToken)
		log.Println("User Authenticated")
		return ctx.Next()
	}

}
