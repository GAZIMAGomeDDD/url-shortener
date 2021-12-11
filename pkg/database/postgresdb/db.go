package postgresdb

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

func NewDB(ctx context.Context, connString string) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("failed to parse conn string (%s): %w", connString, err)
	}

	pool, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	if err = pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping postgres due to error: %w", err)
	}

	return pool, nil
}
