package controllers

import (
	"evermos-backend/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// 🔹 Melihat Semua Log Perubahan Stok
func GetProductLogs(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)

	var logs []models.ProductLog
	// ✅ Preload "Product" dan "Product.Store" agar data toko produk ikut dimuat
	if err := db.Preload("Product.Store").Find(&logs).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch product logs"})
	}

	return c.JSON(logs)
}

