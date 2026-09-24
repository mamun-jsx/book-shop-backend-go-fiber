package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/mamun-jsx/book-shop-backend-go-fiber/internal/app"
	"github.com/mamun-jsx/book-shop-backend-go-fiber/internal/router"
)

func main() {

	application := app.BootstrapApp()
	fiberApp := fiber.New()
	router.SetupRoutes(fiberApp, application)
	serverAddr := fmt.Sprintf(":%s", application.Config.APPport)
	log.Printf("🚀 Server is running on http://localhost%s", serverAddr)
	if err := fiberApp.Listen(serverAddr); err != nil {
		log.Fatalf("server failed to run : %v", err)
	}
}
