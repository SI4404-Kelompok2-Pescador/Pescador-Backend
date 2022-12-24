package dto

type StoreRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Password string `json:"password"`
	OwnerID  string `json:"owner_id"`
}

type StoreLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type StoreResponse struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

type StoreLoginResponse struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	OwnerID string `json:"owner_id"`
	Type    string `json:"type"`
}
