package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       string `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Image    string `json:"image"`
	Password string `json:"password"`
	Type     string `json:"type"`
}

func (u *User) BeforeCreate(_ *gorm.DB) (err error) {
	u.ID = uuid.NewString()
	return
}
