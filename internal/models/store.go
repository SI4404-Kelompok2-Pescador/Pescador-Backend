package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Store struct {
	ID       string `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Password string `json:"password"`
	OwnerID  string `json:"user_id"`
	Owner    User   `json:"user" gorm:"foreignKey:OwnerID"`
	Type     string `json:"type"`
}

func (s *Store) BeforeCreate(tx *gorm.DB) (err error) {
	s.ID = uuid.NewString()
	return
}

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
	Type	string `json:"type"`
}
