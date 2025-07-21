package database

import (
	"blogapp/models"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// ✅ Load .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("⚠️  Warning: .env file not found or failed to load")
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("❌ MySQL Connection Failed:", err)
		panic("Failed to connect to MySQL")
	}

	db.AutoMigrate(&models.Post{})
	DB = db
}
