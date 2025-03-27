package redis

import (
	"github.com/redis/go-redis/v9"
	"context"
)

var Client *redis.Client

func InitRedis() {
	Client = redis.NewClient(&redis.Options{
		Addr:     "reg-redis-2zftop:6379",
		Password: "reg2025", // no password set
		DB:       0,         // use default DB
	})
}

func SetUserData(token string, data interface{}) error {
	ctx := context.Background()
	return Client.Set(ctx, token, data, 0).Err() // Store indefinitely
}

func GetUserData(token string) (string, error) {
	ctx := context.Background()
	return Client.Get(ctx, token).Result()
} 