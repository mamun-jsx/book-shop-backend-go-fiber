package database

import (
	"fmt"
	"log"

	"github.com/mamun-jsx/book-shop-backend-go-fiber/config"
	"github.com/mamun-jsx/book-shop-backend-go-fiber/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// database connection
func Connect(cfg *config.Config) (*gorm.DB, error) {

	var dltor gorm.Dialector

	// if env does not have any nenon db database
	if cfg.DATABASEurl != "" {
		dltor = postgres.Open(cfg.DATABASEurl)

	} else {
		log.Fatal("Can not find database string as NEON db")
	}
	db, err := gorm.Open(dltor, &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("Can not open database connection")
	}
	err = db.AutoMigrate(&models.Book{})
	if err != nil {
		return nil, fmt.Errorf("Unable to migrate the database table")
	}
	log.Println("🚀 Neon PostgreSQL Connection table successfully migrate ")
	return db, nil
}
