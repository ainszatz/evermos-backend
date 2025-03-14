package controllers

import (
	"fmt"
	"evermos-backend/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// 🔹 Membuat Transaksi (Membeli Produk)
func CreateTransaction(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	userID := c.Locals("user_id").(uint)

	transaction := new(models.Transaction)
	if err := c.BodyParser(transaction); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	// 🔹 Cek apakah user memiliki toko
	var user models.User
	if err := db.Preload("Store").First(&user, userID).Error; err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "User not found"})
	}

	// 🔹 Cek apakah produk tersedia
	var product models.Product
	if err := db.Preload("Store").First(&product, transaction.ProductID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Product not found"})
	}

	// 🔹 Cek stok produk cukup
	if product.Stock < transaction.Quantity {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Not enough stock available"})
	}

	transaction.TotalPrice = float64(transaction.Quantity) * product.Price
	transaction.UserID = userID

	// 🔹 Simpan transaksi
	if err := db.Create(&transaction).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create transaction"})
	}

	// 🔹 Update stok produk
	product.Stock -= transaction.Quantity
	db.Save(&product)

	// 🔹 Catat di ProductLog
	productLog := models.ProductLog{
		ProductID: product.ID,
		Change:    -transaction.Quantity,
		Note:      fmt.Sprintf("Purchase by User %d", userID),
	}
	db.Create(&productLog)

	// ✅ **Preload semua data yang dibutuhkan**
	db.Preload("User.Store").Preload("Product.Store").First(&transaction, transaction.ID)

	return c.Status(fiber.StatusCreated).JSON(transaction)
}


// 🔹 Melihat Semua Transaksi User
func GetTransactions(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	userID := c.Locals("user_id").(uint)

	var transactions []models.Transaction
	// ✅ Preload semua relasi yang diperlukan
	if err := db.Preload("User.Store").Preload("Product.Store").
		Where("user_id = ?", userID).Find(&transactions).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch transactions"})
	}

	return c.JSON(transactions)
}



