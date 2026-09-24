package service

import (
	"errors"

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

// create book

func (s *BookService) CreateBook(book *models.Book) error {
	if book.Price < 0 {
		return errors.New("Book price must be grater then 0")
	}
	// return s.repo.Create(book)
	return s.repo.Create(book)
}

// get all books
func (s *BookService) GetAllBooks() ([]models.Book, error) {
	return s.repo.FindAll()
}

// get a single book by id
func (s *BookService) GetBookByID(id uuid.UUID) (*models.Book, error) {
	return s.repo.FindByID(id)
}

// update a single book

func (s *BookService) UpdateBook(book *models.Book) error {
	return s.repo.Update(book)
}

// delete a single books by id
func (s *BookService) DeleteBookByID(id uuid.UUID) error {
	return s.repo.Delete(id)
}
