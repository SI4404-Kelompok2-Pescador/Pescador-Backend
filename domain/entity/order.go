package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	ID             string    `json:"id" gorm:"primary_key, type:uid, default:uuid_generate_v4()"`
	UserID         string    `json:"user_id"`
	User           User      `json:"user" gorm:"foreignKey:UserID"`
	ShippingMethod string    `json:"shipping_method"` // "JNE", "J&T", "TIKI", "POS", "GOJEK", "GRAB"
	ShippingPrice  float64   `json:"shipping_price"`
	TotalPrice     float64   `json:"total_price"`
	Status         string    `json:"status"` // "Pending", "On Process", "On Delivery", "Delivered"
	StoreID        string    `json:"store_id"`
	Store          Store     `json:"store" gorm:"foreignKey:StoreID"`
	CreatedAt      time.Time `json:"created_at"`
}

func (order *Order) BeforeCreate(tx *gorm.DB) (err error) {
	order.ID = uuid.NewString()
	return
}
