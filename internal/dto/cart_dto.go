package dto

type CartRequest struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

type CartResponse struct {
	ID        string  `json:"id"`
	ProductName string  `json:"product_name"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

type CartListResponse struct {
	ID          string  `json:"id"`
	ProductName string  `json:"product_name"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
}
