package handlers

import (
	"github.com/Maksim-Gol/neuralService/internal/models"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	app.Get("/api", welcome)
	app.Post("/calls", StoreCall)
}
func StoreCall(ctx *fiber.Ctx) error {
	var callData models.ServiceCall

	if err := ctx.BodyParser(&callData); err != nil {
		// ? Как мне здесь получить доступ к логгеру, который я задал в мейне?
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid JSON"})

	}

	return ctx.JSON(fiber.Map{"message": "success", "data": callData})
}

func welcome(c *fiber.Ctx) error {
	return c.SendString("I am a sent string")
}
