package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/mamun-jsx/book-shop-backend-go-fiber/internal/models"
	"github.com/mamun-jsx/book-shop-backend-go-fiber/internal/service"
	"gorm.io/gorm"
)

type BookHandler struct {
	service *service.BookService
}

func NewBookHandler(service *service.BookService) *BookHandler {
	return &BookHandler{service: service}
}

// CreateBook creates a single book
func (h *BookHandler) CreateBook(c fiber.Ctx) error {
	book := new(models.Book)

	if err := c.Bind().Body(book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.service.CreateBook(book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(book)
}

// get all books

func (h *BookHandler) GetAllBooks(c fiber.Ctx) error {
	// 1. ALWAYS check the database query error first!
	books, err := h.service.GetAllBooks()

	if len(books) == 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Currently We have no books Listed"})
	}
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not fetch the books"})
	}
	return c.JSON(books)
}

// delete a single book by id

func (h *BookHandler) DeleteBookByID(c fiber.Ctx) error {
	idStr := c.Params("id")
	parsedID, err := uuid.Parse(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid UUID format"})
	}

	if err := h.service.DeleteBookByID(parsedID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not delete book"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Book successfully deleted"})
}
