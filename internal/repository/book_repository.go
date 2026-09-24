package repository

import (
	"github.com/mamun-jsx/book-shop-backend-go-fiber/internal/models"
	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}

// create book to database
func (r *BookRepository) Create(book *models.Book) error {
	return r.db.Create(book).Error
}
