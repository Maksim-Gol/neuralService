package main

import (
	"github.com/Maksim-Gol/neuralService/handlers"
	"github.com/gofiber/fiber/v2"
)


func main(){
	app := fiber.New()

	handlers.RegisterRoutes(app)
	app.Listen(":3000")
}