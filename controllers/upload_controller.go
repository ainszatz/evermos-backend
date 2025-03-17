package controllers

import (
	"evermos-backend/config"
	"evermos-backend/models"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"errors"

	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
)

const UploadPath = "./uploads/"

// 🔹 Upload gambar produk
func UploadProductImage(c *fiber.Ctx) error {
	// Ambil ID produk dari parameter URL
	productID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid product ID"})
	}

	// Ambil user_id dari Locals
	userIDRaw := c.Locals("user_id")
	if userIDRaw == nil {
		fmt.Println("user_id is missing in Locals")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "User not authenticated"})
	}

	userID, ok := userIDRaw.(uint)
	if !ok {
		fmt.Println("Failed to parse user_id")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to parse user ID"})
	}

	fmt.Println("User ID from Locals:", userID) // Debugging

	// Ambil produk dari database, termasuk informasi toko
	var product models.Product
	if err := config.DB.Preload("Store").First(&product, productID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Product not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database error"})
	}

	// Ambil informasi toko berdasarkan StoreID produk
	var store models.Store
	if err := config.DB.First(&store, product.StoreID).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Store not found"})
	}

	// Pastikan user yang login adalah pemilik toko
	if store.UserID != userID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "You are not allowed to upload images for this product"})
	}

	// Ambil file dari request
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "No image file uploaded"})
	}

	// Tentukan path penyimpanan file
	savePath := filepath.Join("uploads", fmt.Sprintf("product_%d_%s", productID, file.Filename))

	// Simpan file ke folder
	if err := c.SaveFile(file, savePath); err != nil {
		log.Println("Error saving file:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save image"})
	}

	// Simpan URL gambar ke database
	product.ImageURL = savePath
	if err := config.DB.Save(&product).Error; err != nil {
		os.Remove(savePath) // Hapus file jika gagal menyimpan data
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update product image"})
	}

	// Berikan respons sukses
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":   "Image uploaded successfully",
		"image_url": savePath,
	})
}



// 🔹 Upload avatar user
func UploadUserAvatar(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	userID := c.Locals("user_id").(uint)

	file, err := c.FormFile("avatar")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "No avatar uploaded"})
	}

	filename := fmt.Sprintf("user_%d%s", userID, filepath.Ext(file.Filename))
	filePath := UploadPath + filename

	if err := c.SaveFile(file, filePath); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save avatar"})
	}

	// Update URL avatar di database
	var user models.User
	if err := db.First(&user, userID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	user.Avatar = filePath
	db.Save(&user)

	return c.JSON(fiber.Map{"message": "Avatar uploaded successfully", "avatar_url": filePath})
}


// 🔹 Upload logo toko
func UploadStoreLogo(c *fiber.Ctx) error {
    // Validasi DB
    dbInterface := c.Locals("db")
    if dbInterface == nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database connection not found"})
    }

    db, ok := dbInterface.(*gorm.DB)
    if !ok {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get database connection"})
    }

    storeID := c.Params("id")

    // Ambil user_id dari context
    userIDInterface := c.Locals("user_id")
    if userIDInterface == nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "User not authenticated"})
    }

    userID, ok := userIDInterface.(uint)
    if !ok {
        fmt.Println("Failed to parse user_id:", userIDInterface)
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to parse user ID"})
    }

    // Ambil user dari database
    var user models.User
    if err := db.First(&user, userID).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
    }

    fmt.Println("User found:", user.ID) // Debugging user

    // Ambil data toko dari database
    var store models.Store
    if err := db.First(&store, storeID).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Store not found"})
    }

    // Cek apakah user adalah pemilik toko
    if store.UserID != user.ID {
        return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "You are not allowed to change this store's logo"})
    }

    // Ambil file dari request
    file, err := c.FormFile("logo")
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "No logo uploaded"})
    }

    // Tentukan nama file dan lokasi penyimpanan
    filename := fmt.Sprintf("store_%s%s", storeID, filepath.Ext(file.Filename))
    filePath := UploadPath + filename

    // Simpan file
    if err := c.SaveFile(file, filePath); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save logo"})
    }

    // Update URL logo di database
    store.LogoURL = filePath
    db.Save(&store)

    return c.JSON(fiber.Map{"message": "Store logo uploaded successfully", "logo_url": filePath})
}



