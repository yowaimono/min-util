package minutil

import (
	"context"


	"github.com/go-redis/redis/v8"
)

// RedisTokenStore 是基于 Redis 的 Token 存储实现
type RedisTokenStore struct {
	client *redis.Client
}

// NewRedisTokenStore 创建一个新的 RedisTokenStore
func NewRedisTokenStore(client *redis.Client) *RedisTokenStore {
	return &RedisTokenStore{
		client: client,
	}
}

// Get 获取 Token 的状态
func (r *RedisTokenStore) Get(token string) (bool, error) {
	ctx := context.Background()
	val, err := r.client.Get(ctx, token).Result()
	if err == redis.Nil {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return val == "1", nil
}

// Set 设置 Token 的状态
func (r *RedisTokenStore) Set(token string, invalid bool) error {
	ctx := context.Background()
	val := "0"
	if invalid {
		val = "1"
	}
	return r.client.Set(ctx, token, val, 0).Err()
}