package model

import (
	"errors"
	"time"
)

type Account struct {
	ID        int64     `json:"id"`
	OwnerName string    `json:"owner_name"`
	Balance   float64   `json:"balance"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"created_at"`
}

func NewAccount(ownerName string, balance float64, currency string) (*Account, error) {
	if ownerName == "" {
		return nil, errors.New("owner name is empty")
	}
	if balance < 0 {
		return nil, errors.New("balance cannot be negative")
	}
	if currency == "" {
		return nil, errors.New("currency required")
	}

	return &Account{
		OwnerName: ownerName,
		Balance:   balance,
		Currency:  currency,
	}, nil
}
