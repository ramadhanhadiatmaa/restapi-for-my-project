package main

import (
	"apicsmfib/controller/usercontroller"
	"apicsmfib/models"

	"github.com/gofiber/fiber/v2"
)

func main() {
	models.ConnectionDatabase()

	app := fiber.New()

	api := app.Group("/api")
	user := api.Group("/users")

	user.Get("/", usercontroller.Index)
	user.Get("/:username", usercontroller.Show)
	user.Post("/", usercontroller.Create)
	user.Put("/:username", usercontroller.Update)
	user.Delete("/:username", usercontroller.Delete)

	app.Listen(":8300")
}