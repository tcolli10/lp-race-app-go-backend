package main

import (
	"example/lp-race-app-go-backend/handlers"

	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {

	//	"/"
	app.Get("/", handlers.ListUsers)

	//	"/user"
	app.Post("/user", handlers.CreateUser)
	app.Delete("/user", handlers.DeleteUserByName)
}
