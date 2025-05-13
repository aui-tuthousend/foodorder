package main

import (
    "foodorder/internal/database"
    "foodorder/internal/auth"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/logger"
    "github.com/gofiber/swagger"
    _ "foodorder/docs" // for Swagger docs
)

//	@title			Food Ordering API
//	@version		1.0
//	@description	API untuk aplikasi pemesanan makanan online
//	@host			localhost:3000
//	@BasePath		/api
func main() {
    app := fiber.New()
    app.Use(logger.New())

    // Connect DB
    database.Connect()

    // Auto migrate model
    db := database.DB
    db.AutoMigrate(&auth.User{})

    // Swagger
    app.Get("/swagger/*", swagger.HandlerDefault)

    // Register routes
    api := app.Group("/api")
    auth.RegisterRoutes(api, db)
    // food.RegisterRoutes(api, db)

    app.Listen(":3000")
}
