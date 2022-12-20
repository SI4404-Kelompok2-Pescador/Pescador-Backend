package user

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
	Password string `json:"password"`
	Type     string `json:"type"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewString()
	return
}

type UserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Password string `json:"password"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	Type    string `json:"type"`
}

// type UserToken struct {
// 	ID     string `json:"id" gorm:"primary_key, type:uid, default:uuid_generate_v4()"`
// 	UserID string `json:"user_id"`
// 	User   User   `json:"user" gorm:"foreignKey:UserID"`
// 	Type   string `json:"type"`
// 	Token  string `json:"token"`
// }

// func (u *UserToken) BeforeCreate(tx *gorm.DB) (err error) {
// 	u.ID = uuid.NewString()
// 	return
// }
