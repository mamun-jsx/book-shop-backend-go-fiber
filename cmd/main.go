package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/mamun-jsx/book-shop-backend-go-fiber/config"
	"github.com/mamun-jsx/book-shop-backend-go-fiber/internal/database"
)

func main() {

	// load config and database
	cfg := config.LoadEnv()

	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal("failed to connect with database", err)
	}
	fmt.Println(db)

	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	if err := app.Listen(":8080"); err != nil {
		log.Fatal("Server failed to Run")
	}
}
