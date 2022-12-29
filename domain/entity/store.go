package entity

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
	Image    string `json:"image"`
	Password string `json:"password"`
	OwnerID  string `json:"user_id"`
	Owner    User   `json:"user" gorm:"foreignKey:OwnerID"`
	Type     string `json:"type"`
}

func (s *Store) BeforeCreate(_ *gorm.DB) (err error) {
	s.ID = uuid.NewString()
	return
}
