package psql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kaibling/iggy/pkg/config"
)

func New(ctx context.Context, cfg config.DBConfig) (*pgxpool.Pool, error) {
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable",
		cfg.DBUser,
		cfg.DBDatabase,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort)

	return pgxpool.New(ctx, connStr)
}
