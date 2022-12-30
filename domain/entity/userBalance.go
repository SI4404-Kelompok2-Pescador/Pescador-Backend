package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserBalance struct {
	ID      string  `json:"id" gorm:"primary_key, type:uid, default:uuid_generate_v4()"`
	UserID  string  `json:"user_id"`
	User    User    `json:"user" gorm:"foreignKey:UserID"`
	Balance float64 `json:"balance" default:"0"`
}

func (ub *UserBalance) BeforeCreate(_ *gorm.DB) (err error) {
	ub.ID = uuid.NewString()
	return
}
