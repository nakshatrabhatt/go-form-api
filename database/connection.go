package database

import (
	"fmt"
	"log"
	"os"
	
	"github.com/nakshatrabhatt/go-form-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDB initializes the database connection
func ConnectDB() {
	var err error
	
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	
	fmt.Println("Connected to PostgreSQL database!")
	
	// Auto migrate only the models that exist
	err = DB.AutoMigrate(
		&models.User{},
		&models.Form{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}