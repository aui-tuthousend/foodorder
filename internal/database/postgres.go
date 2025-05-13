package database

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "os"
    "log"
    
    "github.com/joho/godotenv"
)

var DB *gorm.DB

func Connect() *gorm.DB {
    err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
    dsn := os.Getenv("DATABASE_URL")
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    DB = db
    return db
}