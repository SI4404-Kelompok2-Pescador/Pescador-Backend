package dto

import (
	"github.com/golang-jwt/jwt/v4"
)

type CustomClaims struct {
	UserName string `json:"user_name"`
	jwt.RegisteredClaims
}
