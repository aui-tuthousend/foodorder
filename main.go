package main

import (
	"foodorder/database"
	"foodorder/routes"
	"foodorder/docs"
	e "foodorder/entities"
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
        docs.SwaggerInfo.Host = "foodorder-production-057c.up.railway.app" // ganti dengan domain production kamu
    }
    
    app := fiber.New(fiber.Config{
		EnablePrintRoutes: true,
	})
    app.Use(logger.New())

    database.Connect()

    db := database.DB
    db.AutoMigrate(&e.User{})
    // db.AutoMigrate(&e.User{}, &e.otherEntities{})

    routes.SetupRoutes(app, db)
    app.Get("/swagger/*", swagger.HandlerDefault)

    app.Listen(":8080")
}
