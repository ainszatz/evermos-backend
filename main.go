package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"evermos-backend/config"
	"evermos-backend/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDB()

	db := config.DB
	if db == nil {
		log.Fatal("❌ Database connection is nil!")
	}

	app := fiber.New(fiber.Config{})
	app.Use(func(c *fiber.Ctx) error {
		fmt.Println("✅ Database connection stored in Locals")
		c.Locals("db", db)
		return c.Next()
	})
	

	// Register Routes
	routes.AuthRoutes(app)
	routes.UserRoutes(app)
	routes.StoreRoutes(app)
	routes.AddressRoutes(app)
	routes.CategoryRoutes(app)
	routes.ProductRoutes(app)
	routes.TransactionRoutes(app)
	routes.ProductLogRoutes(app)


	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Welcome to Evermos Backend API"})
	})

	port := os.Getenv("APP_PORT")
	fmt.Println("Server running on port", port)
	log.Fatal(app.Listen(":" + port))
}
