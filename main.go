package main

import (
    e "foodorder/entities"
    "foodorder/database"
    "foodorder/features/auth/common"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/logger"
    "github.com/gofiber/swagger"
    _ "foodorder/docs" // for Swagger docs
)

//	@title			Food Ordering API
//	@version		1.0
//	@description	API untuk aplikasi pemesanan makanan online
//	@host			127.0.0.1:8080
//	@BasePath		/

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
    app := fiber.New()
    app.Use(logger.New())

    database.Connect()

    db := database.DB
    db.AutoMigrate(&e.User{})

    app.Get("/swagger/*", swagger.HandlerDefault)

    api := app.Group("/api")
    common.RegisterRoutes(api, db)
    // food.RegisterRoutes(api, db)

    app.Listen(":8080")
}
