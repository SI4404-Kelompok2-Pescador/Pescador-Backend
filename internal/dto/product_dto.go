package dto

type ProductRequest struct {
	Name         string  `json:"name"`
	Price        float64 `json:"price"`
	Stock        int     `json:"stock"`
	Description  string  `json:"description"`
	Image        string  `json:"image"`
	CategoryName string  `json:"category_name"`
	StoreID      string  `json:"store_id"`
}

type ProductResponse struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Category    string  `json:"category"`
	StoreName   string  `json:"store_name"`
}
