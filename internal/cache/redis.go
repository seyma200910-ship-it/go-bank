package cahce

import (
	"context"
	"encoding/json"
	"fmt"
	"service/internal/model"
	"time"

	"github.com/redis/go-redis/v9"
)

type AccountRedis struct {
	client *redis.Client
}

func NewAccountRedis(client *redis.Client) *AccountRedis {
	return &AccountRedis{
		client: client,
	}
}

func accountKey(id int) string {
	return fmt.Sprintf("account:%d", id)
}

func (c *AccountRedis) GetByID(ctx context.Context, id int) (*model.Account, error) {
	value, err := c.client.Get(ctx, accountKey(id)).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil // cache miss
		}
		return nil, err
	}

	var acc *model.Account
	if err := json.Unmarshal([]byte(value), &acc); err != nil {
		return nil, err
	}

	return acc, nil
}

func (c *AccountRedis) Set(ctx context.Context, acc *model.Account, ttl time.Duration) error {
	data, err := json.Marshal(acc)
	if err != nil {
		return err
	}

	return c.client.Set(ctx, accountKey(acc.ID), data, ttl).Err()
}

func (c *AccountRedis) Delete(ctx context.Context, id int) error {
	return c.client.Del(ctx, accountKey(id)).Err()
}
