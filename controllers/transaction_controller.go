package controllers

import (
	"fmt"
	"time"
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

	// ✅ Ambil query parameter untuk pagination
	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 10)
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * limit

	// ✅ Ambil query parameter untuk filtering
	productID := c.QueryInt("product_id")
	startDate := c.Query("start_date") // Format: YYYY-MM-DD
	endDate := c.Query("end_date")     // Format: YYYY-MM-DD

	var transactions []models.Transaction
	var total int64

	// 🔹 Query dasar
	query := db.Model(&models.Transaction{}).Where("user_id = ?", userID)

	// 🔹 Tambahkan filter jika ada query parameter
	if productID > 0 {
		query = query.Where("product_id = ?", productID)
	}
	if startDate != "" {
		parsedStartDate, err := time.Parse("2006-01-02", startDate)
		if err == nil {
			query = query.Where("created_at >= ?", parsedStartDate)
		}
	}
	if endDate != "" {
		parsedEndDate, err := time.Parse("2006-01-02", endDate)
		if err == nil {
			parsedEndDate = parsedEndDate.Add(24 * time.Hour) // Tambahkan satu hari agar filter mencakup seluruh hari
			query = query.Where("created_at < ?", parsedEndDate)
		}
	}

	// ✅ Hitung total transaksi sesuai filter
	query.Count(&total)

	// ✅ Ambil data dengan filter, pagination, dan preload
	if err := query.Preload("User.Store").Preload("Product.Store").
		Limit(limit).Offset(offset).
		Find(&transactions).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch transactions"})
	}

	// ✅ Response dengan metadata pagination
	return c.JSON(fiber.Map{
		"total":       total,
		"page":        page,
		"limit":       limit,
		"transactions": transactions,
	})
}





