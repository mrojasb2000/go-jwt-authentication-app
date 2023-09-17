package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mrojasb2000/go-jwt-authentication-app/database"
	"github.com/mrojasb2000/go-jwt-authentication-app/routes"
)

func main() {
	database.Connect()

	app := fiber.New()

	routes.Setup(app)

	app.Listen(":3000")
}
