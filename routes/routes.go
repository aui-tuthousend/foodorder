package routes

import (
    "github.com/gofiber/fiber/v2"
    "gorm.io/gorm"

    auth "foodorder/features/auth/common"
    user "foodorder/features/user/common"
    // "foodorder/internal/food"
    // "foodorder/internal/order"
    // "foodorder/middleware"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
    api := app.Group("/api")

    auth.RegisterRoutes(api, db)
    user.RegisterRoutes(api, db)
    
}
