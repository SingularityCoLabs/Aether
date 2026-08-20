// Package database owns PostgreSQL lifecycle, migrations, and health checks.
package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Database wraps the shared pgx pool.
type Database struct {
	pool *pgxpool.Pool
}

// Open creates and verifies a PostgreSQL pool.
func Open(ctx context.Context, databaseURL string) (*Database, error) {
	config, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		return nil, fmt.Errorf("parse database URL: %w", err)
	}
	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("create database pool: %w", err)
	}
	database := &Database{pool: pool}
	if err := database.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("ping database: %w", err)
	}
	return database, nil
}

// Ping implements the server readiness health-check contract.
func (d *Database) Ping(ctx context.Context) error {
	return d.pool.Ping(ctx)
}

// Pool exposes pgx to storage modules without exposing it to transports.
func (d *Database) Pool() *pgxpool.Pool {
	return d.pool
}

func (d *Database) Close() {
	d.pool.Close()
}
