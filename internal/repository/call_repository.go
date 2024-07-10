package repository

import (
	"context"
	"log/slog"

	"github.com/Maksim-Gol/neuralService/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
	"fmt"
)

type Repository struct {
	DB *pgxpool.Pool
}

func (r *Repository) SaveCall(ctx context.Context, callData models.ServiceCall) (string, error) {
	query := `INSERT INTO service_calls VALUES 
		($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.DB.Exec(context.Background(), query,
		callData.UserID,
		callData.ModelID,
		callData.RequestID,
		callData.Cost,
		callData.Status,
		callData.CallTime,
		callData.Metadata)
	if err != nil {
		fmt.Printf("QueryRow in store calls gave an error: %+v\n", err)
	}
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
