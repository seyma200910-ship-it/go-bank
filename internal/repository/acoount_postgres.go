package repository

import (
	"context"
	"fmt"
	"service/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AccountPostgres struct {
	db *pgxpool.Pool
}

func NewAccountPostgres(db *pgxpool.Pool) *AccountPostgres {
	return &AccountPostgres{
		db: db,
	}
}

func (a *AccountPostgres) Create(ctx context.Context, acc *model.Account) (*model.Account, error) {
	query := `
		INSERT INTO accounts (owner_name, balance, currency)
		VALUES ($1, $2, $3)
		RETURNING id, owner_name, balance, currency, created_at
	`

	err := a.db.QueryRow(ctx, query, acc.OwnerName, acc.Balance, acc.Currency).Scan(
		&acc.ID,
		&acc.OwnerName,
		&acc.Balance,
		&acc.Currency,
		&acc.CreatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}

	return acc, nil
}
