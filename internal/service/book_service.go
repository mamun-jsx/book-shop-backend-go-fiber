package service

import (
	"errors"
	"strings"

	"github.com/google/uuid"
	"github.com/mamun-jsx/book-shop-backend-go-fiber/internal/models"
	"github.com/mamun-jsx/book-shop-backend-go-fiber/internal/repository"
)

type BookService struct {
	repo *repository.BookRepository
}

func NewBookService(repo *repository.BookRepository) *BookService {
	return &BookService{repo: repo}
}

// CreateBook validates and creates a new book
func (s *BookService) CreateBook(book *models.Book) error {
	if strings.TrimSpace(book.Name) == "" {
		return errors.New("book name is required")
	}
	if strings.TrimSpace(book.AuthorName) == "" {
		return errors.New("author name is required")
	}
	if book.Price < 0 {
		return errors.New("book price cannot be negative")
	}
	if book.Pages < 0 {
		return errors.New("book pages cannot be negative")
	}
	return s.repo.Create(book)
}

// GetAllBooks retrieves all books
func (s *BookService) GetAllBooks() ([]models.Book, error) {
	return s.repo.FindAll()
}

// GetBookByID retrieves a single book by ID
func (s *BookService) GetBookByID(id uuid.UUID) (*models.Book, error) {
	return s.repo.FindByID(id)
}

// UpdateBook validates and updates a book
func (s *BookService) UpdateBook(book *models.Book) error {
	if strings.TrimSpace(book.Name) == "" {
		return errors.New("book name is required")
	}
	if strings.TrimSpace(book.AuthorName) == "" {
		return errors.New("author name is required")
	}
	if book.Price < 0 {
		return errors.New("book price cannot be negative")
	}
	if book.Pages < 0 {
		return errors.New("book pages cannot be negative")
	}
	return s.repo.Update(book)
}

// DeleteBookByID deletes a single book by ID
func (s *BookService) DeleteBookByID(id uuid.UUID) error {
	return s.repo.Delete(id)
}
