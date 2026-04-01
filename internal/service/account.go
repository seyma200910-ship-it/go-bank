package service

import (
	"context"
	"service/internal/model"
	"time"
)

type AccountCache interface {
	GetByID(ctx context.Context, id int) (*model.Account, error)
	Set(ctx context.Context, user *model.Account, ttl time.Duration) error
	Delete(ctx context.Context, id int) error
}

type AccountRepository interface {
	GetByID(ctx context.Context, id int) (*model.Account, error)
	Create(ctx context.Context, acc *model.Account) (*model.Account, error)
}

type AccountService struct {
	repo  AccountRepository
	cache AccountCache
}

func NewAccountService(repo AccountRepository, cache AccountCache) *AccountService {
	return &AccountService{
		repo:  repo,
		cache: cache,
	}
}

func (a *AccountService) CreateAccount(ctx context.Context, owner string, balance float64, email string, currency string) (*model.Account, error) {
	acc, err := model.NewAccount(owner, balance, email, currency)
	if err != nil {
		return nil, err
	}

	return a.repo.Create(ctx, acc)

}

func (a *AccountService) GetByID(ctx context.Context, id int) (*model.Account, error) {

	acc, err := a.cache.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if acc != nil {
		return acc, nil
	}

	acc, err = a.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if err := a.cache.Set(ctx, acc, 300*time.Second); err != nil {
		return nil, err
	}

	return acc, nil
}
