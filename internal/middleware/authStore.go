package middleware

import (
	"Pescador-Backend/config"
	"Pescador-Backend/domain/entity"
	"log"

	"github.com/gofiber/fiber/v2"
)

func AuthStore(c Config) fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		header := ctx.GetReqHeaders()

		if _, ok := header["Authorization"]; !ok {
			return c.Unauthorized(ctx)
		}

		storeToken := entity.StoreToken{}
		err := config.DB.Where("token = ?", header["Authorization"]).First(&storeToken).Error
		if err != nil {
			return c.Unauthorized(ctx)
		}

		if storeToken.Type != "store" {
			return c.Unauthorized(ctx)
		}

		ctx.Locals("store", storeToken)
		log.Println("Store Authenticated")
		return ctx.Next()
	}

}
