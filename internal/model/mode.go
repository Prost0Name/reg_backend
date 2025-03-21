package model

import (
	"backend/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Login    string `gorm:"uniqueIndex"`
	Password string
}

var DB *gorm.DB

func InitDatabase(dsn config.DSNConfig) {
	var err error
	dsnString := "host=" + dsn.Host + " user=" + dsn.User + " password=" + dsn.Pass + " dbname=" + dsn.DBName + " port=" + dsn.Port + " sslmode=" + dsn.SSLMode
	DB, err = gorm.Open(postgres.Open(dsnString), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB.AutoMigrate(&User{})
}
