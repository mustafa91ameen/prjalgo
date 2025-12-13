package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(ctx context.Context, dsn string, maxRetries int, retryDelay time.Duration) (*pgxpool.Pool, error) {
	var pool *pgxpool.Pool
	var err error

	for attempt := 0; attempt < maxRetries; attempt++ {
		// Try to create connection pool
		pool, err = pgxpool.New(ctx, dsn)
		if err != nil {
			if attempt < maxRetries-1 {
				delay := retryDelay * time.Duration(1<<uint(attempt)) // Exponential backoff
				log.Printf("Database connection attempt %d/%d failed: %v. Retrying in %v...", attempt+1, maxRetries, err, delay)
				time.Sleep(delay)
				continue
			}
			return nil, fmt.Errorf("failed to create connection pool after %d attempts: %w", maxRetries, err)
		}

		// Try to ping database
		if err = pool.Ping(ctx); err != nil {
			pool.Close() // Close failed connection
			if attempt < maxRetries-1 {
				delay := retryDelay * time.Duration(1<<uint(attempt)) // Exponential backoff
				log.Printf("Database ping attempt %d/%d failed: %v. Retrying in %v...", attempt+1, maxRetries, err, delay)
				time.Sleep(delay)
				continue
			}
			return nil, fmt.Errorf("failed to ping database after %d attempts: %w", maxRetries, err)
		}

		// Success!
		if attempt > 0 {
			log.Printf("Database connection established after %d attempts", attempt+1)
		}
		return pool, nil
	}

	return nil, fmt.Errorf("failed to connect to database after %d attempts", maxRetries)
}
