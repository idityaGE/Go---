package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/idityaGE/go-mongo-gofiber/controllers"
)

func Handler(app fiber.Router) {
	app.Get("/u", controllers.GetAllUsers)
	app.Get("/u/:id", controllers.GetUserById)
	app.Post("/u", controllers.CreateUser)
	app.Put("/u/:id", controllers.UpdateUser)
	app.Delete("/u/:id", controllers.DeleteUser)
}
