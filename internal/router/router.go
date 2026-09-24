package router

import (
	"github.com/gofiber/fiber/v3"
	"github.com/mamun-jsx/book-shop-backend-go-fiber/internal/app"
)

func SetupRoutes(fiberApp *fiber.App, application *app.App) {
	fiberApp.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Welcome to Book Shop APi")
	})

	api := fiberApp.Group("/api/v1")
	_ = api
}
