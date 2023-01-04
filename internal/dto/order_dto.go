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
	Status         string    `json:"status"`
	CreatedAt      time.Time `json:"created_at"`
}

type UpdateOrderRequest struct {
	Status string `json:"status"`
}

type StoreOrderResponse struct {
	ID             string    `json:"id"`
	ShippingMethod string    `json:"shipping_method"`
	ShippingPrice  float64   `json:"shipping_price"`
	TotalPrice     float64   `json:"total_price"`
	UserName       string    `json:"user_name"`
	UserAddress    string    `json:"user_address"`
	UserPhone      string    `json:"user_phone"`
	UserEmail      string    `json:"user_email"`
	Status         string    `json:"status"`
	CreatedAt      time.Time `json:"created_at"`
}
