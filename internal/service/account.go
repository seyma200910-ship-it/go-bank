package service

import (
	"service/internal/model"
	"service/internal/repository"
)

type AccountService struct {
	repo *repository.AccountRepository
}

func NewAccountService(repo *repository.AccountRepository) *AccountService {
	return &AccountService{
		repo: repo,
	}
}

func (a *AccountService) CreateAccount(owner string, balance float64, currency string) (*model.Account, error) {
	acc, err := model.NewAccount(owner, balance, currency)
	if err != nil {
		return nil, err
	}

	return a.repo.Create(acc)

}
