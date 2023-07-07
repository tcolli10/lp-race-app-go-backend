package main

import (
	"example/lp-race-app-go-backend/db"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db.Connect()
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}
