package routes

import (
	"github.com/ChitrangGoyani/task-mgmt-auth/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/user", controllers.GetUser)
	api.Post("/register", controllers.Register)
	api.Post("/login", controllers.Login)
	api.Post("/logout", controllers.Logout)
}
