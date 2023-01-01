package dto

import "time"

type OrderRequest struct {
	ShippingMethod string `json:"shipping_method"`
}

type OrderResponse struct {
	ID             string    `json:"id"`
	ShippingMethod string    `json:"shipping_method"`
	ShippingPrice  float64   `json:"shipping_price"`
	TotalPrice     float64   `json:"total_price"`
	CreatedAt      time.Time `json:"created_at"`
}
