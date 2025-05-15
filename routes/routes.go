package routes

import (
    "github.com/gofiber/fiber/v2"
    "gorm.io/gorm"

    auth "foodorder/features/auth/common"
    // "foodorder/internal/food"
    // "foodorder/internal/order"
    // "foodorder/middleware"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
    api := app.Group("/api")

    auth.RegisterRoutes(api, db)
    // food.RegisterRoutes(api.Group("/foods"), db)
    // order.RegisterRoutes(api.Group("/orders"), db, middleware.JWTProtected())
}
