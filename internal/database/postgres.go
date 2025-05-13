package database

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "os"
    "log"
)

var DB *gorm.DB

func Connect() *gorm.DB {
    dsn := os.Getenv("DATABASE_URL")
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    DB = db
    return db
}