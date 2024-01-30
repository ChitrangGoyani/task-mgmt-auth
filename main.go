package main

import (
	"github.com/ChitrangGoyani/task-mgmt-auth/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	// connect to database
	database.Connect()
	// setup cors
	app.Use(cors.New(
		cors.Config{
			AllowCredentials: true, // this allows cookies in each request back and forth
		},
	))
	// call routes function
	app.Listen(":8000")
}
