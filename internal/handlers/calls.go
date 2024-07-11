package handlers

import (
	//"github.com/Maksim-Gol/neuralService/internal/repository"
	"context"
	"log/slog"
	"github.com/Maksim-Gol/neuralService/internal/models"
	"github.com/gofiber/fiber/v2"
)

type RepositoryProvider interface {
	SaveCall(ctx context.Context, call models.ServiceCall) error
	GetCalls(ctx context.Context, user_id string, model_id string) ([]models.ServiceCall, error)
}

func RegisterRoutes(app *fiber.App, db RepositoryProvider) {
	app.Get("/calls", GetCall(db))
	app.Post("/calls", StoreCall(db))
}

func StoreCall(db RepositoryProvider) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var callData models.ServiceCall
		if err := ctx.BodyParser(&callData); err != nil {
			slog.Debug("Erorr parsing json body", err)
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid JSON"})

		}
		slog.Info("Stored call into database")
		err := db.SaveCall(context.Background(), callData)
		if err != nil {
			slog.Debug("Error while storing call into database", err)
		}
		return ctx.JSON(fiber.Map{"message": "success", "data": callData})

	}
}

func GetCall(db RepositoryProvider) fiber.Handler {
	return func(ctx *fiber.Ctx) error{
	user, model := ctx.Query("user","Genadiy"), ctx.Query("model", "alphafold")
	slog.Info("Accepted Get-request for user ", user, "and model ", model, ".")
	calls, err := db.GetCalls(context.Background(), user, model)
	if err != nil {
		slog.Debug("Error getting calls from database")
	}
	return ctx.JSON(fiber.Map{"message": "success", "data": calls})
	}
}
