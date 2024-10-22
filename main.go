package main

import (
	"github.com/taufiq-azr/ecommerce-go-api/config"
	"github.com/taufiq-azr/ecommerce-go-api/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	config.ConnectDB()

	app := fiber.New()

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":8080"))
}
