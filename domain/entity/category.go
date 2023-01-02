package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	ID   string `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	Name string `json:"name" gorm:"primaryKey"`
}

func (c *Category) BeforeCreate(_ *gorm.DB) (err error) {
	c.ID = uuid.NewString()
	return
}
