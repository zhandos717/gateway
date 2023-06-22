package database

import (
	"github.com/zhandos717/gateway/app/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	// Устанавливаем соединение с базой данных SQLite
	db, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&models.IpAddres{})
	if err != nil {
		return err
	}

	DB = db

	return nil
}
