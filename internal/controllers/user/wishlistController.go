package user

import (
	"Pescador-Backend/config"
	"Pescador-Backend/domain/entity"
	"net/http"
	"time"

	"Pescador-Backend/internal/dto"

	"github.com/gofiber/fiber/v2"
)

func (u *UserImplementation) AddWishlist(c *fiber.Ctx) error {
	req := dto.WishlistRequest{}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	var product entity.Product

	err := config.DB.Where("id = ?", req.ProductID).First(&product).Error

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Product not found",
		})
	}

	user := c.Locals("user").(entity.UserToken)

	var wishlist entity.Wishlist

	// if product already in wishlist, delete wishlist

	// get productID in user wishlist

	err = config.DB.Where("user_id = ?", user.UserID).Where("product_id = ?", req.ProductID).First(&wishlist).Error

	if err == nil {
		// if product in wishlist, delete wishlist

		err = config.DB.Delete(&wishlist).Error

		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"message": "Delete wishlist failed",
			})
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"message": "Delete wishlist success",
		})
	}

	// if product not in wishlist, create new wishlist

	wishlist = entity.Wishlist{
		UserID:    user.UserID,
		ProductID: req.ProductID,
		CreatedAt: time.Now(),
	}

	err = config.DB.Create(&wishlist).Error

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Create wishlist failed",
		})
	}

	response := dto.WishlistResponse{
		ID:           wishlist.ID,
		ProductName:  product.Name,
		ProductPrice: product.Price,
		ProductImage: product.Image,
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Create wishlist success",
		"data":    response,
	})

}

func (u *UserImplementation) ShowWishlist(c *fiber.Ctx) error {

	user := c.Locals("user").(entity.UserToken)

	var wishlist []entity.Wishlist

	err := config.DB.Where("user_id = ?", user.UserID).Find(&wishlist).Error

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Get wishlist failed",
		})
	}

	var response []dto.WishlistResponse

	for _, v := range wishlist {
		var product entity.Product

		err = config.DB.Where("id = ?", v.ProductID).First(&product).Error

		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"message": "Get wishlist failed",
			})
		}

		response = append(response, dto.WishlistResponse{
			ID:           v.ID,
			ProductName:  product.Name,
			ProductPrice: product.Price,
			ProductImage: product.Image,
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Get wishlist success",
		"data":    response,
	})
}
