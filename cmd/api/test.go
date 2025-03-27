package main

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func main() {
	// Создаем контекст
	ctx := context.Background()

	// Используем IP-адрес вместо имени хоста
	opt, err := redis.ParseURL("redis://default:reg2025@87.242.100.33:6379")
	if err != nil {
		log.Fatalf("Ошибка парсинга URL: %v", err)
	}

	rdb := redis.NewClient(opt)

	// Тестируем соединение
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Ошибка подключения к Redis: %v", err)
	}
	fmt.Println("Успешное подключение:", pong)

	// Записываем тестовые данные
	err = rdb.Set(ctx, "test_key", "Hello, Redis!", 0).Err()
	if err != nil {
		log.Fatalf("Ошибка записи: %v", err)
	}
	fmt.Println("Ключ test_key записан!")

	// Читаем записанное значение
	val, err := rdb.Get(ctx, "test_key").Result()
	if err != nil {
		log.Fatalf("Ошибка чтения: %v", err)
	}
	fmt.Println("Прочитанное значение:", val)
}
