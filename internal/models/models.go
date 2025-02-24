package models

import (
	"backend/internal/config"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Login    string `gorm:"uniqueIndex"`
	Password string
}

var DB *gorm.DB

func InitDatabase() {
	var err error
	DB, err = gorm.Open(postgres.Open(config.DSN), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	DB.AutoMigrate(&User{})
}