package repository

import (
	"context"
	"errors"
	"fmt"
	"service/internal/model"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ExistsEmail = errors.New("Аккаунт с таким email уже существует")

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
		INSERT INTO accounts (owner_name, balance, email,currency)
		VALUES ($1, $2, $3, $4)
		RETURNING id, owner_name, balance,email, currency, created_at
	`

	err := a.db.QueryRow(ctx, query, acc.OwnerName, acc.Balance, acc.Email, acc.Currency).Scan(
		&acc.ID,
		&acc.OwnerName,
		&acc.Balance,
		&acc.Email,
		&acc.Currency,
		&acc.CreatedAt,
	)

	if err != nil {
		if pgErr, ok := errors.AsType[*pgconn.PgError](err); ok {
			if pgErr.Code == "23505" && pgErr.ConstraintName == "accounts_email_key" {
				return nil, ExistsEmail
			}
		}
		return nil, fmt.Errorf("create account: %w", err)
	}

	return acc, nil
}

func (a *AccountPostgres) GetByEmail(ctx context.Context, email string) (bool, error) {
	var exists bool

	query := `
	SELECT EXISTS (
		SELECT 1 FROM accounts WHERE email = $1
	)
	`

	if err := a.db.QueryRow(ctx, query, email).Scan(&exists); err != nil {
		return false, fmt.Errorf("ошибка чтения: %w", err)
	}

	return exists, nil
}

func (a *AccountPostgres) GetByID(ctx context.Context, id int) (*model.Account, error) {
	acc := &model.Account{}

	query := `
	SELECT id, owner_name, balance, email, currency, created_at
	FROM accounts
	WHERE id = $1
	`

	if err := a.db.QueryRow(ctx, query, id).Scan(
		&acc.ID,
		&acc.OwnerName,
		&acc.Balance,
		&acc.Email,
		&acc.Currency,
		&acc.CreatedAt); err != nil {
		return nil, fmt.Errorf("ошибка чтения: %w", err)
	}

	return acc, nil
}
