package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"evermos-backend/config"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to Database
	config.ConnectDB()

	// Setup Fiber App
	app := fiber.New()

	// Define Root Route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Welcome to Evermos Backend API",
		})
	})

	// Start Server
	port := os.Getenv("APP_PORT")
	fmt.Println("Server running on port", port)
	log.Fatal(app.Listen(":" + port))
}
