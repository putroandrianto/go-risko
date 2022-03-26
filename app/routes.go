package app

import (
	"go-risko/controllers"

	"github.com/gofiber/fiber/v2"
)

func router() *fiber.App {
	app := fiber.New()

	app.Get("/", controllers.Index)
	app.Get("/health", controllers.Health)
	app.Get("/users", controllers.Users)
	app.Get("/users/:id", controllers.User)
	app.Post("/users", controllers.CreateUser)
	app.Put("/users/:id", controllers.UpdateUser)
	app.Delete("/users/:id", controllers.DeleteUser)

	return app
}