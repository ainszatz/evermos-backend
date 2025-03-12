package controllers

import (
	"evermos-backend/config"
	"evermos-backend/middleware"
	"evermos-backend/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// Register User
func Register(c *fiber.Ctx) error {
	type RegisterInput struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Phone    string `json:"phone"`
		Password string `json:"password"`
	}

	var input RegisterInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Cek apakah email atau phone sudah ada
	var existingUser models.User
	if err := config.DB.Where("email = ? OR phone = ?", input.Email, input.Phone).First(&existingUser).Error; err == nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Email atau No HP sudah terdaftar"})
	}

	// Hash password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	// Buat user baru
	user := models.User{Name: input.Name, Email: input.Email, Phone: input.Phone, Password: string(hashedPassword)}
	config.DB.Create(&user)

	// Buat toko otomatis untuk user baru
	store := models.Store{UserID: user.ID, Name: input.Name + "'s Store"}
	config.DB.Create(&store)

	return c.Status(http.StatusCreated).JSON(fiber.Map{"message": "User registered successfully!"})
}

// Login User
func Login(c *fiber.Ctx) error {
	type LoginInput struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var input LoginInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	var user models.User
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid email or password"})
	}

	// Cek password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid email or password"})
	}

	// Ambil role terbaru dari database untuk generate token
	var updatedUser models.User
	config.DB.First(&updatedUser, user.ID)

	// Generate JWT Token
	token, err := middleware.GenerateToken(user.ID, user.Role)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate token"})
	}

	return c.JSON(fiber.Map{"token": token})
}
