package main

import (
	"apicsmfib/models"
	"apicsmfib/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	models.ConnectionDatabase()

	app := fiber.New()

	routes.Route(app)

	app.Listen(":8080")
}