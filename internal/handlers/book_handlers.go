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

// GetAllBooks retrieves all books
func (h *BookHandler) GetAllBooks(c fiber.Ctx) error {
	books, err := h.service.GetAllBooks()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not fetch books"})
	}

	if books == nil {
		books = []models.Book{}
	}

	return c.Status(fiber.StatusOK).JSON(books)
}

// GetBookByID retrieves a single book by ID
func (h *BookHandler) GetBookByID(c fiber.Ctx) error {
	idStr := c.Params("id")
	parsedID, err := uuid.Parse(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid UUID format"})
	}

	book, err := h.service.GetBookByID(parsedID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not fetch book"})
	}

	return c.Status(fiber.StatusOK).JSON(book)
}

// UpdateBook updates a single book by ID
func (h *BookHandler) UpdateBook(c fiber.Ctx) error {
	idStr := c.Params("id")
	parsedID, err := uuid.Parse(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid UUID format"})
	}

	existingBook, err := h.service.GetBookByID(parsedID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Book not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not fetch book"})
	}

	var updateData models.Book
	if err := c.Bind().Body(&updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if updateData.Name != "" {
		existingBook.Name = updateData.Name
	}
	if updateData.AuthorName != "" {
		existingBook.AuthorName = updateData.AuthorName
	}
	if updateData.Pages > 0 {
		existingBook.Pages = updateData.Pages
	}
	if updateData.Type != "" {
		existingBook.Type = updateData.Type
	}
	if updateData.Price >= 0 {
		existingBook.Price = updateData.Price
	}
	if updateData.ShortDescription != "" {
		existingBook.ShortDescription = updateData.ShortDescription
	}

	if err := h.service.UpdateBook(existingBook); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(existingBook)
}

// DeleteBookByID deletes a single book by ID
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
