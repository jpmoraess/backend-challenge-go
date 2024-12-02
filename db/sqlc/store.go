package db

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

// Store store provides all functions to execute db queries and transactions
type Store interface {
	Querier
}

// SQLStore provides all functions to execute SQL queries and transactions
type SQLStore struct {
	*Queries
	pool *pgxpool.Pool
}

// NewStore creates a new store
func NewStore(pool *pgxpool.Pool) *SQLStore {
	return &SQLStore{
		pool:    pool,
		Queries: New(pool),
	}
}
