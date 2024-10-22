package main

import (
	"ecommerce-api/config"
	"ecommerce-api/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	config.ConnectDB()

	app := fiber.New()

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":8080"))
}
