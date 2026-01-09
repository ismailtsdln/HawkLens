package db

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	Client *redis.Client
}

func NewRedisClient(addr string, password string, db int) *RedisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	return &RedisClient{Client: rdb}
}

func (r *RedisClient) SetCache(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return r.Client.Set(ctx, key, data, expiration).Err()
}

func (r *RedisClient) GetCache(ctx context.Context, key string, target interface{}) error {
	val, err := r.Client.Get(ctx, key).Result()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(val), target)
}

func (r *RedisClient) CheckRateLimit(ctx context.Context, key string, limit int, window time.Duration) (bool, error) {
	count, err := r.Client.Incr(ctx, key).Result()
	if err != nil {
		return false, err
	}
	if count == 1 {
		r.Client.Expire(ctx, key, window)
	}
	return int(count) <= limit, nil
}

func (r *RedisClient) Ping(ctx context.Context) error {
	return r.Client.Ping(ctx).Err()
}
