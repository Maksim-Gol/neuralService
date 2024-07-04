package main

import (
	"github.com/Maksim-Gol/neuralService/handlers"
	"github.com/Maksim-Gol/neuralService/internal/config"
	"github.com/gofiber/fiber/v2"
	"log/slog"
	"os"
)


const (
	envLocal = "local"
	envProd = "dev"
	envDev = "prod"
)

func main(){
	// TODO: config clenaenv
	cfg := config.MustLoad()
	// TODO: logger slog
	log := setupLogger(cfg.Env)
	log.Info("Start neuralService", slog.String("env", cfg.Env))
	log.Debug("Debug messages are enabled")
	app := fiber.New()

	handlers.RegisterRoutes(app)
	app.Listen(cfg.Port)
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)

	}
	return log

}