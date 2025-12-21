package storage 

import (
	"testing"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

const testDatabaseURL = "postgres://notify:notify@localhost:5433/notify_test?sslmode=disable"

func setupTestStorage(t *testing.T) *PostgresStorage {
	ctx := context.Background()

	pool, err := pgxpool.New(ctx, testDatabaseURL)
	if err != nil {
		t.Fatalf("failed to connect to test database: %v", err)
	}

	cleanAllTables(t, pool)

	return NewPostgresStorage(pool)
}

func cleanAllTables(t *testing.T, pool *pgxpool.Pool) {
	ctx := context.Background()
    
    _, err := pool.Exec(ctx, "DELETE FROM templates")
    if err != nil {
        t.Fatalf("failed to clean templates: %v", err)
    }
    
    _, err = pool.Exec(ctx, "DELETE FROM preferences")
    if err != nil {
        t.Fatalf("failed to clean preferences: %v", err)
    }

	_, err = pool.Exec(ctx, "DELETE FROM delivery_logs")
    if err != nil {
        t.Fatalf("failed to clean delivery_logs: %v", err)
    }
}