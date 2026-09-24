package service

import (
	"errors"

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