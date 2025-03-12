package controllers

import (
	"evermos-backend/config"
	"evermos-backend/models"
	// "net/http"

	"github.com/gofiber/fiber/v2"
)

// 🏡 Buat Address Baru
func CreateAddress(c *fiber.Ctx) error {
    // Ambil user_id dari middleware JWT
    userID, ok := c.Locals("user_id").(uint)
    if !ok {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
    }

    // Ambil data request
    address := new(models.Address)
    if err := c.BodyParser(address); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
    }

    // Set user_id sebelum insert ke database
    address.UserID = userID

    // Simpan ke database
    if err := config.DB.Create(&address).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create address"})
    }

    return c.Status(fiber.StatusCreated).JSON(address)
}


// 📍 Ambil Semua Address Milik User yang Sedang Login
func GetAddressesByUser(c *fiber.Ctx) error {
	// Ambil `user_id` dari JWT (pastikan middleware sudah mengatur `Locals`)
	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	var addresses []models.Address
	if err := config.DB.Where("user_id = ?", userID).Find(&addresses).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(addresses)
}

// ✏️ Update Address (dengan validasi kepemilikan)
func UpdateAddress(c *fiber.Ctx) error {
	id := c.Params("id")
	var address models.Address

	// Cek apakah address dengan ID ini ada
	if err := config.DB.First(&address, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Address not found"})
	}

	// Validasi bahwa address ini milik user yang sedang login
	userID, ok := c.Locals("user_id").(uint)
	if !ok || address.UserID != userID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Unauthorized access"})
	}

	// Parse request baru ke dalam struct sementara
	var updateData models.Address
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Update hanya field yang diperbolehkan
	address.Street = updateData.Street
	address.City = updateData.City
	address.ZipCode = updateData.ZipCode

	// Simpan perubahan
	if err := config.DB.Save(&address).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(address)
}

// 🗑️ Hapus Address (dengan validasi kepemilikan)
func DeleteAddress(c *fiber.Ctx) error {
	id := c.Params("id")
	var address models.Address

	// Cek apakah address ada
	if err := config.DB.First(&address, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Address not found"})
	}

	// Validasi kepemilikan
	userID, ok := c.Locals("user_id").(uint)
	if !ok || address.UserID != userID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Unauthorized access"})
	}

	// Hapus Address
	if err := config.DB.Delete(&address).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Address deleted successfully"})
}
