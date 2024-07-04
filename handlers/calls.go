package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App){
	app.Get("/api", welcome)
}


func welcome (c *fiber.Ctx) error {
	return c.SendString("I am a string")
}