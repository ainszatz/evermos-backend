package tests

import (
	"fmt"
	"os"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"evermos-backend/middleware"
)

func TestGenerateToken(t *testing.T) {
	userID := uint(1) // Example user ID
	role := "user"    // Example role

	// Set the JWT_SECRET environment variable for testing
	os.Setenv("JWT_SECRET", "supersecretkey")

	token, err := middleware.GenerateToken(userID, role)
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	fmt.Println("Generated Token:", token)

	// Validate the token
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !parsedToken.Valid {
		t.Fatalf("Invalid token: %v", err)
	}
}
