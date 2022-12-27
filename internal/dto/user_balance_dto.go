package dto

type UserBalanceRequest struct {
	UserID  string  `json:"user_id"`
	Balance float64 `json:"balance"`
}

type UserBalanceResponse struct {
	Balance float64 `json:"balance"`
}
