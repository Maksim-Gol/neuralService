package repository

import (
	"context"
	"github.com/Maksim-Gol/neuralService/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
)

type Repository struct {
	DB *pgxpool.Pool
}

func InitDB(connectionString string, log *slog.Logger) (*Repository, error) {
	var err error
	var dbPool *pgxpool.Pool
	dbPool, err = pgxpool.New(context.Background(), connectionString)
	if err != nil {
		//slog.Error("Unable to connect to database", "error", err)
		return nil, err
	}
	_, err = dbPool.Acquire(context.Background())
	if err != nil {
		//slog.Error("Unable to acquire connection to database", "error", err)
		return nil, err
	}
	slog.Info("Connected to postgres")
	return &Repository{DB: dbPool}, nil
}

func (r *Repository) SaveCall(ctx context.Context, callData models.ServiceCall) error {
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
		slog.Error("QueryRow in store calls gave an error", "error", err)
		return err
	}
	return nil
}

func (r *Repository) GetCalls(ctx context.Context, user string, model string) ([]models.ServiceCall, error) {
	var calls []models.ServiceCall
	query := `SELECT * FROM service_calls WHERE user_id = $1 AND model_id = $2`
	rows, err := r.DB.Query(ctx, query, user, model)
	if err != nil {
		slog.Error("Error querying values from db", "error", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var call models.ServiceCall
		err := rows.Scan(&call.UserID, &call.ModelID, &call.RequestID, &call.Cost,
			&call.Status, &call.CallTime, &call.Metadata)
		if err != nil {
			slog.Error("Error scanning values from row to ServiceCall object", "error", err)
			return nil, err
		}
		calls = append(calls, call)

	}
	if ctx.Err() != nil {
		slog.Error("Context error", "error", ctx.Err())
		return nil, ctx.Err()
	}
	if rows.Err() != nil {
		slog.Error("Rows iteration error", "error", rows.Err())
		return nil, rows.Err()
	}
	return calls, nil

}
