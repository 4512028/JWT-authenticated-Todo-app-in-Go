package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// For development: load .env if exists
	_ = os.Getenv("ENV") // optional: use this to control loading

	// if os.Getenv("ENV") != "production" {
	// 	_ = godotenv.Load() // useful in local dev only
	// }
	log.Println("DB_HOST =", os.Getenv("DB_HOST"))

	// Read from env
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port,
	)

	log.Printf("Connecting with DSN: %s\n", dsn) // DEBUG

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	log.Println("✅ Connected to PostgreSQL database.")
}
