package database

import (
	"log"

	"github.com/bergsantana/go-contacts/internal/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSQLiteDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("contacts.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	if err := db.AutoMigrate(&entity.Contact{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	return db
}
