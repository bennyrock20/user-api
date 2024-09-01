package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"taxi-service/config"
)

var DB *gorm.DB

func InitDatabase(cfg config.Config) {
	var err error
	DB, err = gorm.Open(postgres.Open(cfg.DatabaseDSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
}
