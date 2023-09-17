package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mrojasb2000/go-jwt-authentication-app/controllers"
)

func Setup(app *fiber.App) {
	app.Get("/", controllers.Hello)
}
