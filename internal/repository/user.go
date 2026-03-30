package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type User interface {
}

type user struct {
	pool *pgxpool.Pool
}

func NewUser(pool *pgxpool.Pool) User {
	return &user{pool: pool}
}
