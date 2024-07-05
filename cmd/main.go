package main

import (
	"github.com/Maksim-Gol/neuralService/internal/config"
	"github.com/Maksim-Gol/neuralService/internal/handlers"
	"github.com/gofiber/fiber/v2"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envProd  = "dev"
	envDev   = "prod"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	log.Info("Start neuralService", slog.String("env", cfg.Env))
	log.Debug("Debug messages are enabled")

	app := fiber.New()
	handlers.RegisterRoutes(app)

	app.Listen(cfg.HTTP.Port)
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	default:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log

}
