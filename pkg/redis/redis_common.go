package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

var (
	ctx    = context.Background()
	Client *redis.Client
)

func NewRedisClient() *redis.Client {
	Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return Client
}

func SetExpiration(key string, val any, exp time.Duration) error {
	return Client.Set(ctx, key, val, exp).Err()
}

func Set(key string, val any) error {
	return SetExpiration(key, val, 0)
}

func Get(key string) (string, error) {
	return Client.Get(ctx, key).Result()
}
