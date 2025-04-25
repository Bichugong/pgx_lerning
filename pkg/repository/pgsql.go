package repository

import (
	"context"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PGRepo struct {
	mu sync.Mutex
	pool *pgxpool.Pool
}

func New(connStr string) (*PGRepo, error) {
	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		return nil,  err
	}
	return &PGRepo{
		pool: pool,
	}, nil
}