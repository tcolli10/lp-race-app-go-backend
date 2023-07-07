package handlers

import (
	"example/lp-race-app-go-backend/db"
	"example/lp-race-app-go-backend/models"

	"github.com/gofiber/fiber/v2"
)

func ListUsers(ctx *fiber.Ctx) error {
	users := []models.User{}

	db.DB.Db.Find(&users)
	return ctx.Status(200).JSON(users)
}
