package database

import (
	"log"

	"github.com/pressly/goose"
	"gorm.io/gorm"
)

func ApplyMigrations(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Error getting SQL connection:", err)
	}
	if err := goose.Up(sqlDB, "migrations"); err != nil {
		log.Fatal("Error applying migrations:", err)
	}
	log.Println("Migrations applied successfully")
}
