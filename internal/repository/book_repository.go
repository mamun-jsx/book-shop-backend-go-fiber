package repository

import (
	"github.com/google/uuid"
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

// get all books
func (r *BookRepository) FindAll() ([]models.Book, error) {
	var books []models.Book
	err := r.db.Find(&books).Error
	return books, err
}

// get a single book by ID

func (r *BookRepository) FindByID(id uuid.UUID) (*models.Book, error) {
	var book models.Book
	err := r.db.First(&book, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}

// Update a single book into Database

func (r *BookRepository) Update(book *models.Book) error {
	return r.db.Save(book).Error
}

// delete a single book by id

func (r *BookRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Book{}, "id = ?", id).Error
}
