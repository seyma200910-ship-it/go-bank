CREATE TABLE accounts (
    id BIGSERIAL PRIMARY KEY,
    owner_name TEXT NOT NULL,
    balance NUMERIC(12,2) NOT NULL CHECK (balance >= 0),
    currency TEXT NOT NULL DEFAULT 'RUB' CHECK (currency IN ('USD', 'EUR', 'RUB')),
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);