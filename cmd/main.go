package main

import (
	"fmt"
	"github.com/Maksim-Gol/neuralService/internal/config"
	"github.com/Maksim-Gol/neuralService/internal/handlers"
	"github.com/Maksim-Gol/neuralService/internal/repository"
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

	// Init config, logger
	cfg := config.MustLoad()
	log := setupLogger(cfg.Env)
	slog.SetDefault(log)

	//Connecting to postgres
	DBconnectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.Postgres.DBUser, cfg.Postgres.DBPassword, cfg.Postgres.DBHost,
		cfg.Postgres.DBPort, cfg.Postgres.DBName)

	dbPool := repository.InitDB(DBconnectionString, log)

	//Initial logs
	log.Info("Start neuralService", slog.String("env", cfg.Env))
	log.Debug("Debug messages are enabled")

	//Starting App
	app := fiber.New()
	handlers.RegisterRoutes(app, dbPool)

	app.Listen(cfg.HTTP.Port)
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	// Choosing logger based on environment(local,development,production)
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
