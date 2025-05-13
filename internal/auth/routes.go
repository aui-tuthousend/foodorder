package auth

import (
    "foodorder/middleware"
    "github.com/gofiber/fiber/v2"
    "gorm.io/gorm"
)

func RegisterRoutes(api fiber.Router, db *gorm.DB) {
    group := api.Group("/auth")
    group.Post("/register", Register(db))
    group.Post("/login", Login(db))
    group.Get("/profile", middleware.JWTProtected(), Profile(db))
}
