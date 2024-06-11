package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

// type Database struct {
// 	Conn *sql.DB
// }

// func NewDatabase(db *sql.DB) *Database {
// 	return &Database{Conn: db}
// }

func ConnectDatabase() (*pgxpool.Pool, error) {
	// urlExample := "postgres://username:password@localhost:5432/database_name"

	// Example Keyword/Value
	// user=jack password=secret host=pg.example.com port=5432 dbname=mydb sslmode=verify-ca pool_max_conns=10

	// Example URL
	// postgres://jack:secret@pg.example.com:5432/mydb?sslmode=verify-ca&pool_max_conns=10

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
