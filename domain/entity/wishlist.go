package entity

import (
	"time"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

type Wishlist struct {
	ID        string    `json:"id" gorm:"primary_key, type:uid, default:uuid_generate_v4()"`
	UserID    string    `json:"user_id"`
	User      User      `json:"user" gorm:"foreignKey:UserID"`
	ProductID string    `json:"product_id"`
	Product   Product   `json:"product" gorm:"foreignKey:ProductID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (w *Wishlist) BeforeCreate(_ *gorm.DB) (err error) {
	w.ID = uuid.NewString()
	return
}
