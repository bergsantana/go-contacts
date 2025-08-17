package database

import (
	"log"

	"github.com/bergsantana/go-contacts/internal/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSQLiteDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("db/contacts.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Falha ao conectar ao banco de dados: ", err)
	}

	if err := db.AutoMigrate(&entity.Contact{}); err != nil {
		log.Fatal("Falha ao executar migrations:", err)
	}

	return db
}
