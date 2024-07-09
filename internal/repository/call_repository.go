package repository

import (
	"context"
	"log/slog"

	"github.com/Maksim-Gol/neuralService/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	DB *pgxpool.Pool
}

func (r *Repository) SaveCall(ctx context.Context, call models.ServiceCall) (string, error) {
	return "hm", ctx.Err()
}

func InitDB(connectionString string, log *slog.Logger) *Repository {
	var err error
	var dbPool *pgxpool.Pool
	dbPool, err = pgxpool.New(context.Background(), connectionString)
	if err != nil {
		log.Debug("Unable to connect to database", "error", err)

	}
	return &Repository{DB: dbPool}
}

// func GetDB() *pgxpool.Pool {
// 	return dbPool
// }
