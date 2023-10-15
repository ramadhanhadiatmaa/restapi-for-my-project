package routes

import (
	"apicsmfib/controllers"
	"apicsmfib/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Route(r *fiber.App) {
	api := r.Group("/api")

	api.Get("/users", middlewares.AuthMiddleware, controllers.Index)
	api.Get("/users/:username", middlewares.AuthMiddleware,controllers.Show)
	api.Post("/users", middlewares.AuthMiddleware,controllers.Create)
	api.Put("/users/:username", middlewares.AuthMiddleware,controllers.Update)
	api.Delete("/users/:username", middlewares.AuthMiddleware,controllers.Delete)
}