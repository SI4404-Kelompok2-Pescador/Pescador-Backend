package models

// import (
// 	"Pescador-Backend/internal/models/user"

// 	"github.com/google/uuid"
// 	"gorm.io/gorm"
// )

// type UserToken struct {
// 	ID     string    `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
// 	UserID string    `json:"user_id"`
// 	User   user.User `json:"user" gorm:"foreignKey:UserID"`
// 	Type   string    `json:"type"`
// 	Token  string    `json:"token"`
// }

// func (u *UserToken) BeforeCreate(tx *gorm.DB) (err error) {
// 	u.ID = uuid.NewString()
// 	return
// }
