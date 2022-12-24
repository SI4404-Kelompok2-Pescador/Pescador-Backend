package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserToken struct {
	ID     string `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	UserID string `json:"user_id"`
	User   User   `json:"user" gorm:"foreignKey:UserID"`
	Type   string `json:"type"`
	Token  string `json:"token"`
}

func (u *UserToken) BeforeCreate(_ *gorm.DB) (err error) {
	u.ID = uuid.NewString()
	return
}

type StoreToken struct {
	ID      string `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	StoreID string `json:"store_id"`
	Store   Store  `json:"store" gorm:"foreignKey:StoreID"`
	Type    string `json:"type"`
	Token   string `json:"token"`
}

func (s *StoreToken) BeforeCreate(_ *gorm.DB) (err error) {
	s.ID = uuid.NewString()
	return
}
