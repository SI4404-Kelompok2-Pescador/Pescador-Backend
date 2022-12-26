package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID          string  `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Description string  `json:"description"`
	StoreID     string  `json:"store_id"`
	Store       Store   `json:"store" gorm:"foreignKey:StoreID"`
}

func (p *Product) BeforeCreate(_ *gorm.DB) (err error) {
	p.ID = uuid.NewString()
	return
}
