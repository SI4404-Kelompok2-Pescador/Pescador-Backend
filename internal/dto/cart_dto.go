package dto

type CartRequest struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

type CartResponse struct {
	ID           string  `json:"id"`
	ProductName  string  `json:"product_name"`
	Quantity     int     `json:"quantity"`
	ProductPrice float64 `json:"product_price"`
	TotalPrice   float64 `json:"total_price"`
}

type CartListResponse struct {
	ID           string  `json:"id"`
	ProductName  string  `json:"product_name"`
	Quantity     int     `json:"quantity"`
	ProductPrice float64 `json:"product_price"`
	TotalPrice   float64 `json:"total_price"`
}
