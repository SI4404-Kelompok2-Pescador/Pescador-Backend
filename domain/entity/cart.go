package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Cart struct {
	ID        string `json:"id" gorm:"primary_key, type:uid, default:uuid_generate_v4()"`
	UserID    string `json:"user_id"`
	User      User   `json:"user" gorm:"foreignKey:UserID"`
	ProductID string `json:"product_id"`
	Product   Product `json:"product" gorm:"foreignKey:ProductID"`
	Quantity  int    `json:"quantity"`
}

func (cart *Cart) BeforeCreate(tx *gorm.DB) (err error) {
	cart.ID = uuid.NewString()
	return
}