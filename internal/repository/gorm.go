package repository

import (
	"log"

	"github.com/JamshedJ/REST-api/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBConnection(dsn string) (db *gorm.DB, err error) {
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error connectiong to database", err.Error())
	}

	if err = db.AutoMigrate(&models.Task{}); err != nil {
		log.Fatal("error migrating database")
	}

	return
}