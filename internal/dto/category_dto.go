package dto

type CategoryRequest struct {
	Name string `json:"name"`
}

type DeleteCategoryRequest struct {
	ID string `json:"id"`
}

type CategoryResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
