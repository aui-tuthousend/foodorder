package common

import (
    "foodorder/infrastructure/middleware"
    "foodorder/features/auth/create"
    "foodorder/features/auth/login"
    "foodorder/features/auth/getProfile"
    "github.com/gofiber/fiber/v2"
    "gorm.io/gorm"
)

func RegisterRoutes(api fiber.Router, db *gorm.DB) {
    // if all routes should be authenticated
    // group := api.Group("/routes", middleware.JWTProtected())

    group := api.Group("/auth")
    group.Post("/register", create.Register(db))
    group.Post("/login", login.Login(db))

    // if just one of them
    group.Get("/profile", middleware.JWTProtected(), getprofile.Profile(db))
}
