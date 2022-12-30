package dto

type UserBalanceRequest struct {
	Balance float64 `json:"balance" validate:"required"`
}

type UserBalanceResponse struct {
	Balance float64 `json:"balance"`
}
