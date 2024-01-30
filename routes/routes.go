package routes

import (
	"github.com/ChitrangGoyani/task-mgmt-auth/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Group("/api")
	app.Get("/user", controllers.GetUser)
	app.Post("/register", controllers.Register)
	app.Post("/login", controllers.Login)
	app.Post("/logout", controllers.Logout)
}
