package main

import (
	// "gofiber-rest-api/pkg/database"
	// "gofiber-rest-api/pkg/middleware"
	// "gofiber-rest-api/pkg/routes"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	// Middlewares.
	// middleware.FiberMiddleware(app) // Register Fiber's middleware for app.
	// app.Use(middleware.BodyParserMiddleware())
	// app.Use(middleware.SanitizeMiddleware)

	// Inisialisasi Database
	// database.Connect()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"message": "Hello",
		})
	})

	if os.Getenv("ENV") != "dev" {
		app.Get("/metrics", monitor.New())
	}

	// Inisialisasi Rute
	// routes.RouteInit(app)     // Register a public routes for app.
	// routes.NotFoundRoute(app) // Register route for 404 Error.

	app.Listen(os.Getenv("HOST") + ":" + os.Getenv("PORT"))
}
