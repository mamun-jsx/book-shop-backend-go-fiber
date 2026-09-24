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

// Create inserts a new book record into the database
func (r *BookRepository) Create(book *models.Book) error {
	return r.db.Create(book).Error
}

// FindAll retrieves all books from the database
func (r *BookRepository) FindAll() ([]models.Book, error) {
	var books []models.Book
	err := r.db.Find(&books).Error
	return books, err
}

// FindByID retrieves a single book by ID
func (r *BookRepository) FindByID(id uuid.UUID) (*models.Book, error) {
	var book models.Book
	err := r.db.First(&book, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}

// Update saves changes to an existing book in the database
func (r *BookRepository) Update(book *models.Book) error {
	return r.db.Save(book).Error
}

// Delete removes a single book by ID from the database
func (r *BookRepository) Delete(id uuid.UUID) error {
	result := r.db.Delete(&models.Book{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
