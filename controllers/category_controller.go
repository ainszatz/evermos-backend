package controllers

import (
	"fmt"
	"evermos-backend/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"evermos-backend/config"
)

// CreateCategory (Admin Only)

func CreateCategory(c *fiber.Ctx) error {
	// ✅ Ambil database dari `c.Locals("db")`
	dbRaw := c.Locals("db")
	if dbRaw == nil {
		fmt.Println("❌ Database connection is nil in CreateCategory!")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database connection error"})
	}

	// Konversi `dbRaw` ke *gorm.DB
	db, ok := dbRaw.(*gorm.DB)
	if !ok {
		fmt.Println("❌ Invalid DB type in CreateCategory:", dbRaw)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Invalid database instance"})
	}

	// Parsing request body
	category := new(models.Category)
	if err := c.BodyParser(category); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Simpan kategori ke database
	if err := db.Create(&category).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create category"})
	}

	return c.Status(fiber.StatusCreated).JSON(category)
}

// GetCategories (Public)
func GetCategories(c *fiber.Ctx) error {
	var categories []models.Category
	if err := config.DB.Find(&categories).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch categories"})
	}
	return c.JSON(categories)
}

// UpdateCategory (Admin Only)
func UpdateCategory(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	id := c.Params("id")

	var category models.Category
	if err := db.First(&category, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Category not found"})
	}

	if err := c.BodyParser(&category); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	db.Save(&category)
	return c.JSON(category)
}

// DeleteCategory (Admin Only)
func DeleteCategory(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	id := c.Params("id")

	var category models.Category
	if err := db.First(&category, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Category not found"})
	}

	db.Delete(&category)
	return c.JSON(fiber.Map{"message": "Category deleted"})
}
