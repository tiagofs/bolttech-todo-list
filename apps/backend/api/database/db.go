package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectDatabase() (*pgxpool.Pool, error) {
	// Read from configs, that load from .env file.
	connStr := "user=postgres password=postgres dbname=bolttech host=localhost port=5432 sslmode=disable"
	pool, err := pgxpool.New(context.Background(), connStr)

	if err != nil {
		return nil, err
	}

	// Verify the connection is successful
	if err := pool.Ping(context.Background()); err != nil {
		return nil, err
	}

	return pool, nil
}
