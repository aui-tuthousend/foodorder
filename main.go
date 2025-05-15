package main

import (
	"foodorder/database"
	"foodorder/docs"
	e "foodorder/entities"
	"foodorder/features/auth/common"
	"os"

	_ "foodorder/docs" // for Swagger docs

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
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
    if os.Getenv("ENV") != "production" {
        docs.SwaggerInfo.Host = "127.0.0.1:8080"
    } else {
        docs.SwaggerInfo.Host = "foodorder-aui-tuthousend6429-eedv8bzg.leapcell.dev/" // ganti dengan domain production kamu
    }
    
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
