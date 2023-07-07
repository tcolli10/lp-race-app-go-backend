package handlers

import (
	"example/lp-race-app-go-backend/db"
	"example/lp-race-app-go-backend/models"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(ctx *fiber.Ctx) error {
	user := new(models.User)

	if err := ctx.BodyParser(user); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	db.DB.Db.Create(&user)
	return ctx.Status(200).JSON(user)
}
