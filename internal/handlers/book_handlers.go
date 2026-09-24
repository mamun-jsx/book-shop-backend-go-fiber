package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/mamun-jsx/book-shop-backend-go-fiber/internal/models"
	"github.com/mamun-jsx/book-shop-backend-go-fiber/internal/service"
)

type BookHandler struct {
	service *service.BookService
}

func NewBookHandler(service *service.BookService) *BookHandler {
	return &BookHandler{service: service}
}

func (h *BookHandler) CreateBook(c fiber.Ctx) error {
	book := new(models.Book)

	if err := c.Bind().Body(book); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(book)
}