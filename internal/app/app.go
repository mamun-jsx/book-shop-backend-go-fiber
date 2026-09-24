package app

import (
	"log"

	"github.com/mamun-jsx/book-shop-backend-go-fiber/config"
	"github.com/mamun-jsx/book-shop-backend-go-fiber/internal/database"
	"github.com/mamun-jsx/book-shop-backend-go-fiber/internal/repository"
	"gorm.io/gorm"
)

type App struct {
	Config   *config.Config
	DB       *gorm.DB
	BookRepo *repository.BookRepository
}

func BootstrapApp() *App {
	// load all config
	cfg := config.LoadEnv()
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal("failed to connect with database : %v", err)
	}
	bookRepo := repository.NewBookRepository(db)

	return &App{
		Config:   cfg,
		DB:       db,
		BookRepo: bookRepo,
	}
}
