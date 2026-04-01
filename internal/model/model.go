package model

import (
	"errors"
	"time"
)

type Account struct {
	ID        int       `json:"id"`
	OwnerName string    `json:"owner_name"`
	Balance   float64   `json:"balance"`
	Email     string    `json:"email"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"created_at"`
}

func NewAccount(ownerName string, balance float64, email string, currency string) (*Account, error) {
	if ownerName == "" {
		return nil, errors.New("owner name is empty")
	}

	if email == "" {
		return nil, errors.New("email is empty")
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
		Email:     email,
		Currency:  currency,
	}, nil
}
