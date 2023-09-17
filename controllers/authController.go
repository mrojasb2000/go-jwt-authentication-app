package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Hello(c *fiber.Ctx) error {
	return c.SendString("Hello World!")
}

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return c.JSON(fmt.Sprintf("Error: %s", err))
	}

	return c.JSON(data)
}
