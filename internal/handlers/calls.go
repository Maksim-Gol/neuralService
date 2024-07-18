package handlers

import (
	//"github.com/Maksim-Gol/neuralService/internal/repository"
	"context"
	"log/slog"
	"net/http"
	"github.com/Maksim-Gol/neuralService/internal/models"
	"github.com/gofiber/fiber/v2"
	_ "github.com/Maksim-Gol/neuralService/docs"
)

type RepositoryProvider interface {
	SaveCall(ctx context.Context, call models.ServiceCall) error
	GetCalls(ctx context.Context, user_id string, model_id string) ([]models.ServiceCall, error)
}

func RegisterRoutes(app *fiber.App, db RepositoryProvider, hd func(*fiber.Ctx) error) {
	app.Get("/calls", GetCall(db))
	app.Post("/calls", StoreCall(db))
	app.Get("/swagger/*", hd)
}


// @Summary StoreCall
// @Tags store
// @Description Store call information
// @ID store-call
// @Accept json
// @Produce json
// @Param input body models.ServiceCall true "call info"
// @Success 200 {string} string "ok"
// @Failure 400 "Bad Request"
// @Router /calls [post]
func StoreCall(db RepositoryProvider) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var callData models.ServiceCall
		if err := ctx.BodyParser(&callData); err != nil {
			slog.Debug("Erorr parsing json body", err)
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid JSON"})
		}
		slog.Info("Stored call into database")
		err := db.SaveCall(ctx.Context(), callData)
		if err != nil {
			slog.Debug("Error while storing call into database", err)
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(fiber.Map{"message": "400 Bad Request"})
		}
		return ctx.JSON(fiber.Map{"message": "success", "data": callData})

	}
}

// @Summary GetCall
// @Tags get
// @Description Get call information
// @ID get-call
// @Produce json
// @Param user query string false "User ID" 
// @Param model query string false "Model ID" 
// @Success 200 {object} []models.ServiceCall "success"
// @Failure 400 "Bad Request"
// @Router /calls [get]
func GetCall(db RepositoryProvider) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		user, model := ctx.Query("user", ""), ctx.Query("model", "")
		slog.Info("Accepted Get-request for user ", user, "and model ", model, ".")
		calls, err := db.GetCalls(ctx.Context(), user, model)
		if err != nil {
			slog.Debug("Error getting calls from database")
			return ctx.JSON(fiber.Map{"message": "400 Bad Request"})
		}
		slog.Info("Returning json data:" + calls[0].UserID)
		return ctx.JSON(fiber.Map{"data": calls})
	}
}
