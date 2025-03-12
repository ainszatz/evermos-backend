package controllers

import (
	"evermos-backend/config"
	"evermos-backend/models"
	"github.com/gofiber/fiber/v2"
)

// Get Store Info (Hanya User Login)
func GetStore(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	var store models.Store
	if err := config.DB.Where("user_id = ?", userID).First(&store).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Store not found"})
	}

	return c.JSON(store)
}

// Update Store (Hanya User Login)
func UpdateStore(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	type UpdateStoreInput struct {
		Name string `json:"name"`
	}

	var input UpdateStoreInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	var store models.Store
	if err := config.DB.Where("user_id = ?", userID).First(&store).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Store not found"})
	}

	// Update nama toko
	store.Name = input.Name
	config.DB.Save(&store)

	return c.JSON(fiber.Map{"message": "Store updated successfully"})
}
