package router

import (
	"github.com/gofiber/fiber/v3"
	"github.com/mamun-jsx/book-shop-backend-go-fiber/internal/app"
)

func SetupRoutes(fiberApp *fiber.App, application *app.App) {
	fiberApp.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Welcome to Book Shop API")
	})

	api := fiberApp.Group("/api/v1")

	// ==================| Books API |==================
	api.Post("/books", application.BookHandler.CreateBook)
	api.Post("/book", application.BookHandler.CreateBook)
	api.Get("/books", application.BookHandler.GetAllBooks)
	api.Get("/books/:id", application.BookHandler.GetBookByID)
	api.Put("/books/:id", application.BookHandler.UpdateBook)
	api.Delete("/books/:id", application.BookHandler.DeleteBookByID)
}
