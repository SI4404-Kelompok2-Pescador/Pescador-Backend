package config

import (
	"Pescador-Backend/domain/entity"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Config struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
	SSLMODE  string
}

func Connect(c *Config) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", c.Host, c.User, c.Password, c.DBName, c.Port, c.SSLMODE,
	)

	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	DB = conn

	err = conn.AutoMigrate(
		&entity.User{},
		&entity.Type{},
		&entity.UserType{},
		&entity.StoreType{},
		&entity.UserToken{},
		&entity.StoreToken{},
		&entity.Store{},
		&entity.Product{},
		&entity.UserBalance{},
		&entity.Category{},
	)

	if err != nil {
		log.Fatal(err)
	}

}
