package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Init(databaseURL string) (*pgxpool.Pool, error) {
	fmt.Printf("Initializing database with URL: %s\n", databaseURL)
	return pgxpool.New(context.Background(), databaseURL)
}
