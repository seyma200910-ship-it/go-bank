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
	INSERT INTO account (ownername,balance,currency)
	VALUES ($1,$2,$3)
	RETURNING id, ownername, balance, currency, created_at
	`

	if err := a.db.QueryRow(ctx, query, acc.OwnerName, acc.Balance, acc.Currency).Scan(acc); err != nil {
		return &model.Account{}, fmt.Errorf("create user: %w", err)
	}

	return acc, nil
}
