package main

import (
	"example/lp-race-app-go-backend/db"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db.Connect()
	app := fiber.New()
	setupRoutes(app)

	app.Listen(":3000")
}
