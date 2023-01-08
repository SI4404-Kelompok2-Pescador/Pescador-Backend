package dto

type WishlistRequest struct {
	ProductID string `json:"product_id"`
}

type WishlistResponse struct {
	ID           string  `json:"id"`
	ProductName  string  `json:"product_name"`
	ProductPrice float64 `json:"product_price"`
	ProductImage string  `json:"product_image"`
}
