package service

import (
	"context"
	"service/internal/model"
	"service/internal/repository"
)

type AccountService struct {
	repo *repository.AccountPostgres
}

func NewAccountService(repo *repository.AccountPostgres) *AccountService {
	return &AccountService{
		repo: repo,
	}
}

func (a *AccountService) CreateAccount(ctx context.Context, owner string, balance float64, currency string) (*model.Account, error) {
	acc, err := model.NewAccount(owner, balance, currency)
	if err != nil {
		return nil, err
	}

	return a.repo.Create(ctx, acc)

}
