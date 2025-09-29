package repository

import (
	"TaskFlowAPI/config"
	"TaskFlowAPI/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDB(cfg *config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(cfg.DBPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto migrate
	err = db.AutoMigrate(&models.User{}, &models.Task{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
