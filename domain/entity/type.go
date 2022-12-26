package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Type struct {
	ID   string `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	Name string `json:"name"`
}

func (u *Type) BeforeCreate(_ *gorm.DB) (err error) {
	u.ID = uuid.NewString()
	return
}

type UserType struct {
	ID     string `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	UserID string `json:"user_id"`
	User   User   `json:"user" gorm:"foreignKey:UserID"`
	TypeID string `json:"type_id"`
	Type   Type   `json:"type" gorm:"foreignKey:TypeID"`
}

func (u *UserType) BeforeCreate(_ *gorm.DB) (err error) {
	u.ID = uuid.NewString()
	return
}

type StoreType struct {
	ID      string `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	StoreID string `json:"store_id"`
	Store   Store  `json:"store" gorm:"foreignKey:StoreID"`
	TypeID  string `json:"type_id"`
	Type    Type   `json:"type" gorm:"foreignKey:TypeID"`
}

func (u *StoreType) BeforeCreate(_ *gorm.DB) (err error) {
	u.ID = uuid.NewString()
	return
}
