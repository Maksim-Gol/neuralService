package main

import (
	"fmt"
	_ "github.com/Maksim-Gol/neuralService/docs"
	"github.com/Maksim-Gol/neuralService/internal/config"
	"github.com/Maksim-Gol/neuralService/internal/handlers"
	"github.com/Maksim-Gol/neuralService/internal/repository"
	"github.com/gofiber/fiber/v2" 
	"github.com/gofiber/fiber/v2/middleware/cors"
	//"github.com/gofiber/swagger"
	fiberSwagger "github.com/swaggo/fiber-swagger"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envProd  = "dev"
	envDev   = "prod"
)

// @title NeuralService API
// @version 1.0
// @description API server for NeuralService Application

// @host localhost:3002
// @BasePath /

func main() {

	// Init config, logger
	cfg := config.MustLoad()
	log := setupLogger(cfg.Env)
	slog.SetDefault(log)

	//Connecting to postgres
	DBconnectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.Postgres.DBUser, cfg.Postgres.DBPassword, cfg.Postgres.DBHost,
		cfg.Postgres.DBPort, cfg.Postgres.DBName)

	dbPool, err := repository.InitDB(DBconnectionString, log)
	if err != nil {
		slog.Error("Unable to connect to database", "error", err)
		// ?Мне программу дальше запускать или для прям здесь крашить?
	}

	//Initial logs
	log.Info("Start neuralService", slog.String("env", cfg.Env))
	log.Debug("Debug messages are enabled")

	//Starting App
	app := fiber.New()

	//Allow the server to accept requests from any origin(domain)
	app.Use(cors.New())
	//Potential security issue, 
	//but without it swagger returns TypeError: NetworkError when attempting to fetch resource
	//When executing any http-method

	handlers.RegisterRoutes(app, dbPool, fiberSwagger.WrapHandler)

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
