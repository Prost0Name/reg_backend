package redis

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var Client *redis.Client

func InitRedis() {
	ctx := context.Background()                                              // Создаем контекст
	opt, err := redis.ParseURL("redis://default:reg2025@87.242.100.33:6379") // Используем URL для инициализации
	if err != nil {
		log.Fatalf("Ошибка парсинга URL: %v", err)
	}

	Client = redis.NewClient(opt)

	// Тестируем соединение
	pong, err := Client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Ошибка подключения к Redis: %v", err)
	}
	log.Println("Успешное подключение:", pong)
}

func SetUserData(token string, data interface{}) error {
	ctx := context.Background()
	return Client.Set(ctx, token, data, 0).Err() // Store indefinitely
}

func GetUserData(token string) (string, error) {
	ctx := context.Background()
	return Client.Get(ctx, token).Result()
}
