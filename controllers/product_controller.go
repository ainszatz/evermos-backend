package controllers

import (
	"fmt"
	"evermos-backend/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// ✅ Middleware Auth harus digunakan untuk mendapatkan `user_id`

func CreateProduct(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	userID := c.Locals("user_id").(uint)

	// 🔹 Cek apakah user memiliki toko
	var store models.Store
	if err := db.Where("user_id = ?", userID).First(&store).Error; err != nil {
		fmt.Println("❌ User does not own a store, userID:", userID)
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "User does not own a store"})
	}
	fmt.Println("✅ Store found:", store.ID, "for user:", userID)

	// 🔹 Parsing request body
	product := new(models.Product)
	if err := c.BodyParser(product); err != nil {
		fmt.Println("❌ Failed to parse request body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	fmt.Println("✅ Parsed Product Data:", product)

	// 🔹 Set `store_id` sesuai dengan toko yang dimiliki user
	product.StoreID = store.ID
	fmt.Println("✅ Assigning Store ID:", store.ID, "to Product")

	// 🔹 Simpan produk ke database
	if err := db.Create(&product).Error; err != nil {
		fmt.Println("❌ Failed to create product:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create product"})
	}

	// ✅ **Preload Store setelah insert agar data lengkap**
	db.Preload("Store").First(&product, product.ID)

	return c.Status(fiber.StatusCreated).JSON(product)
}

// 🔹 Lihat Semua Produk (Siapa Saja Bisa Melihat)
func GetProducts(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)

	var products []models.Product
	// ✅ Gunakan Preload agar Store ikut dimuat
	if err := db.Preload("Store").Find(&products).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch products"})
	}

	return c.JSON(products)
}

// 🔹 Update Produk (Hanya Pemilik Toko)
func UpdateProduct(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	userID := c.Locals("user_id").(uint)
	productID := c.Params("id")

	// Cari produk berdasarkan ID
	var product models.Product
	if err := db.First(&product, productID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Product not found"})
	}

	// Pastikan user memiliki toko yang sesuai dengan produk
	var store models.Store
	if err := db.Where("id = ? AND user_id = ?", product.StoreID, userID).First(&store).Error; err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "You do not own this product"})
	}

	// Parsing request body
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Simpan perubahan ke database
	if err := db.Save(&product).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update product"})
	}

	return c.JSON(product)
}

// 🔹 Hapus Produk (Hanya Pemilik Toko)
func DeleteProduct(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	userID := c.Locals("user_id").(uint)
	productID := c.Params("id")

	// Cari produk berdasarkan ID
	var product models.Product
	if err := db.First(&product, productID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Product not found"})
	}

	// Pastikan user memiliki toko yang sesuai dengan produk
	var store models.Store
	if err := db.Where("id = ? AND user_id = ?", product.StoreID, userID).First(&store).Error; err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "You do not own this product"})
	}

	// Hapus dari database
	if err := db.Delete(&product).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete product"})
	}

	return c.JSON(fiber.Map{"message": "Product deleted successfully"})
}
