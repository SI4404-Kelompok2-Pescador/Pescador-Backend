package dto

type ProductRequest struct {
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Description string  `json:"description"`
	Picture     string  `json:"picture"`
	StoreID     string  `json:"store_id"`
}

type ProductResponse struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Description string  `json:"description"`
	Picture     string  `json:"picture"`
	StoreID     string  `json:"store_id"`
}
