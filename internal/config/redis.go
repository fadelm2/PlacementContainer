package config

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

func NewRedis() *redis.Client {
	r := redis.NewClient(&redis.Options{
		Addr:     GetEnv("REDIS_ADDR", "localhost:6379"),
		Password: GetEnv("REDIS_PASSWORD", ""),
		DB:       0,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_ = r.Ping(ctx) // ignore error here, but in prod check
	return r
}
