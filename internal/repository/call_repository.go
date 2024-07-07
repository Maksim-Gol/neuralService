package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
)

var dbPool *pgxpool.Pool

func InitDB(connectionString string, log *slog.Logger) {
	var err error
	dbPool, err = pgxpool.New(context.Background(), connectionString)
	if err != nil {
		log.Debug("Unable to connect to database", "error", err)

	}
}

func GetDB() *pgxpool.Pool {
	return dbPool
}
