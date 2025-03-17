package middleware

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"evermos-backend/models"
	"gorm.io/gorm"
)

func AuthMiddleware(c *fiber.Ctx) error {
    tokenString := c.Get("Authorization")
    fmt.Println("Incoming Token:", tokenString) // Debugging token

    if tokenString == "" {
        fmt.Println("No token provided")
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
    }

    if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
        tokenString = tokenString[7:]
    }

    fmt.Println("JWT_SECRET:", os.Getenv("JWT_SECRET")) // Debugging secret key

    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return []byte(os.Getenv("JWT_SECRET")), nil
    })

    if err != nil || !token.Valid {
        fmt.Println("Token is invalid:", err)
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid Token"})
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        fmt.Println("Invalid token claims")
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid Token Claims"})
    }

    fmt.Println("Claims from JWT:", claims) // Debugging isi token

    userIDRaw, exists := claims["user_id"]
    if !exists {
        fmt.Println("user_id is missing in token claims")
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid Token Claims"})
    }

    // Jika user_id bertipe string di token
    userIDFloat, ok := userIDRaw.(float64)
    if !ok {
        fmt.Println("user_id is not a valid number")
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid user ID format"})
    }

    userID := uint(userIDFloat)
    fmt.Println("user_id set in Locals:", userID) // Debugging sukses

    c.Locals("user_id", userID)

    return c.Next()
}



// Function to generate JWT token
func GenerateToken(userID uint, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Expires in 24 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func AdminOnly() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// ✅ Ambil database dari `c.Locals`
		dbRaw := c.Locals("db")
		if dbRaw == nil {
			fmt.Println("❌ Database connection is nil in AdminOnly Middleware!")
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Database connection error"})
		}

		db, ok := dbRaw.(*gorm.DB)
		if !ok {
			fmt.Println("❌ Invalid DB type in AdminOnly Middleware:", dbRaw)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Invalid database instance"})
		}

		userIDRaw := c.Locals("user_id")
		if userIDRaw == nil {
			fmt.Println("No user_id found in Locals")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized - No user ID"})
		}

		var userID uint
		switch v := userIDRaw.(type) {
		case float64:
			userID = uint(v)
		case int:
			userID = uint(v)
		case uint:
			userID = v
		default:
			fmt.Println("Invalid user_id type:", userIDRaw)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid user ID"})
		}

		var user models.User
		if err := db.First(&user, userID).Error; err != nil {
			fmt.Println("User not found in DB:", userID)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "User not found"})
		}

		if user.Role != "admin" {
			fmt.Println("User is not an admin:", user.Role)
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Access denied"})
		}

		return c.Next()
	}
}

