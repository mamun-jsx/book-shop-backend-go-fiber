package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/mamun-jsx/book-shop-backend-go-fiber/internal/models"
	"github.com/mamun-jsx/book-shop-backend-go-fiber/internal/service"
)

type BookHandler struct {
	service *service.BookService
}

func NewBookHandler(service *service.BookService) *BookHandler {
	return &BookHandler{service: service}
}

// create a single book
func (h *BookHandler) CreateBook(c fiber.Ctx) error {
	book := new(models.Book)

	if err := c.Bind().Body(book); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(book)
}

// get all books
func (h *BookHandler) GetAllBooks(c fiber.Ctx) error {
	// 1. ALWAYS check the database query error first!
	books, err := h.service.GetAllBooks()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not fetch the books from database",
		})
	}

	// 2. Handle the empty state correctly using a 200 OK or 404 status
	if len(books) == 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Currently we have no books listed",
			"data":    []models.Book{}, // Explicitly return an empty list
		})
	}

	// 3. If books exist, return them with 200 OK
	return c.JSON(books)
}

// func (h *BookHandler) GetAllBooks(c fiber.Ctx) error {
// 	books, err := h.service.GetAllBooks()

// 	if len(books) == 0 {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Message": "Currently We have no books Listed"})
// 	}
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not fetch the books"})
// 	}
// 	return c.JSON(books)
// }

// delete a single book by id

func (h *BookHandler) DeleteBookByID(c fiber.Ctx) error {
	isStr := c.Params("id")
	parsedID, err := uuid.Parse(isStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid UUID format"})

	}
	if err := h.service.DeleteBookByID(parsedID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not delete book"})
	}

	return c.JSON(fiber.Map{"message": "Book Successfully deleted"})
}
