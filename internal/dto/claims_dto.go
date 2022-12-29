package dto

import (
	"github.com/golang-jwt/jwt/v4"
)

type CustomClaims struct {
	UserID   string `json:"user_id"`
	UserName string `json:"user_name"`
	jwt.RegisteredClaims
}
