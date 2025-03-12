package controllers

import (
	"evermos-backend/config"
	"evermos-backend/models"
	"github.com/gofiber/fiber/v2"
	// "golang.org/x/crypto/bcrypt"
)

// Get Profile (Hanya User Login)
func GetUserProfile(c *fiber.Ctx) error {
    userID, ok := c.Locals("user_id").(uint)
    if !ok || userID == 0 {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
    }

    var user models.User
    err := config.DB.Preload("Store").First(&user, userID).Error
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
    }

    return c.JSON(user)
}

// Update Profile (Hanya User Login)
func UpdateUserProfile(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	type UpdateUserInput struct {
		Name  string `json:"name"`
		Phone string `json:"phone"`
	}

	var input UpdateUserInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	// Update hanya nama & phone (password tetap aman)
	user.Name = input.Name
	user.Phone = input.Phone
	config.DB.Save(&user)

	return c.JSON(fiber.Map{"message": "Profile updated successfully"})
}

// Delete User (Hanya User Login)
func DeleteUser(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	if err := config.DB.Delete(&models.User{}, userID).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete user"})
	}

	return c.JSON(fiber.Map{"message": "User deleted successfully"})
}
