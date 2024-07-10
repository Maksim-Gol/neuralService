package handlers

import (
	//"github.com/Maksim-Gol/neuralService/internal/repository"
	"context"
	"fmt"
	"github.com/Maksim-Gol/neuralService/internal/models"
	"github.com/gofiber/fiber/v2"
)

type RepositoryProvider interface {
	SaveCall(ctx context.Context, call models.ServiceCall) (string, error)
}

func RegisterRoutes(app *fiber.App, db RepositoryProvider) {
	app.Get("/calls", GetCall)
	app.Post("/calls", StoreCall(db))
}

func StoreCall(db RepositoryProvider) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var callData models.ServiceCall
		if err := ctx.BodyParser(&callData); err != nil {
			// ? Как мне здесь получить доступ к логгеру, который я задал в мейне?
			fmt.Println(fmt.Errorf("%w", err))
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid JSON"})

		}
		db.SaveCall(context.Background(), callData)
		return ctx.JSON(fiber.Map{"message": "success", "data": callData})

	}
}

func GetCall(ctx *fiber.Ctx) error {
	//Getting values from postgres
	/*
		dbPool := repository.GetDB()
		var username string
		err := dbPool.QueryRow(context.Background(), "SELECT * from users;").Scan(&username)
		if err != nil {
			fmt.Println("QueryRow failed", "error", err)
		}
		fmt.Println(username)
	*/
	return ctx.JSON(fiber.Map{"message": "success", "data": "Zhora"})
}
