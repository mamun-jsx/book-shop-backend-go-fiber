package database

import (
	"fmt"
	"log"

	"github.com/mamun-jsx/book-shop-backend-go-fiber/config"
	"github.com/mamun-jsx/book-shop-backend-go-fiber/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connect initializes the PostgreSQL database connection and runs migrations
func Connect(cfg *config.Config) (*gorm.DB, error) {
	if cfg.DATABASEurl == "" {
		return nil, fmt.Errorf("DATABASE_URL is not set")
	}

	db, err := gorm.Open(postgres.Open(cfg.DATABASEurl), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("cannot open database connection: %w", err)
	}

	err = db.AutoMigrate(&models.Book{})
	if err != nil {
		return nil, fmt.Errorf("unable to migrate database tables: %w", err)
	}

	log.Println("🚀 PostgreSQL connection and migration successful")
	return db, nil
}
