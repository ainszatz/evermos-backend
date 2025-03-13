package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
	"evermos-backend/models"
)

var DB *gorm.DB

func ConnectDB() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Panicf("⚠ Error loading .env file: %v", err)
	}

	// Format DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	// Connect ke Database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicf("❌ Failed to connect to database: %v", err)
	}

	// Panggil fungsi untuk migrasi
	MigrateDatabase(db)

	// Simpan koneksi database ke variabel global
	DB = db
	fmt.Println("✅ Database connected and migrated successfully!")
}

// Fungsi untuk menjalankan AutoMigrate
func MigrateDatabase(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.User{}, &models.Store{}, &models.Address{},
		&models.Category{}, &models.Product{}, &models.Transaction{},
		&models.ProductLog{},
	)
	if err != nil {
		log.Panicf("❌ Failed to migrate database: %v", err)
	}
}
