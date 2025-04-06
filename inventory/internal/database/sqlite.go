package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"inventory/internal/model"
)

func InitSQLite() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("inventory.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&model.Product{})
	return db, err
}
