package main

import (
	"context"
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

	//Connecting to postgres
	DBconnectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.Postgres.DBUser, os.Getenv("DBPassword"), cfg.Postgres.DBHost,
		cfg.Postgres.DBPort, cfg.Postgres.DBName)

	repository.InitDB(DBconnectionString, log)
	dbPool := repository.GetDB()

	//Getting values from postgres
	var username string
	err := dbPool.QueryRow(context.Background(), "SELECT * from users;").Scan(&username)
	//? Как выводить эту ошибку и нужно ли это вообще делать
	//В данный момент если что-то с подключением не так(неверный пароль, например),
	//То в логе выдаёт всю строку со всеми данными
	if err != nil {
		log.Debug("QueryRow failed", "error", err)
	}
	fmt.Println(username)

	//Initial logs
	log.Info("Start neuralService", slog.String("env", cfg.Env))
	log.Debug("Debug messages are enabled")

	//Starting App
	app := fiber.New()
	handlers.RegisterRoutes(app)

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
