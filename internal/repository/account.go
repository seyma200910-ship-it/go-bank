package repository

import (
	"service/internal/model"
	"sync"
	"time"
)

type AccountRepository struct {
	mu     sync.Mutex
	data   []model.Account
	nextId int64
}

func NewAccountRepository() *AccountRepository {
	return &AccountRepository{
		data:   []model.Account{},
		nextId: 1,
	}
}

func (a *AccountRepository) Create(acc *model.Account) (*model.Account, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	acc.ID = a.nextId
	acc.CreatedAt = time.Now()

	a.data = append(a.data, *acc)

	a.nextId++

	return acc, nil
}
