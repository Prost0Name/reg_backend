package main

import (
	"backend/internal/app"
	"backend/internal/config"
	"backend/internal/model"
	"backend/internal/redis"
	"fmt"
)

func main() {
	cfg := config.MustLoad()
	redis.InitRedis()
	fmt.Println(cfg.JwtSecret)

	model.InitDatabase(cfg.DSN)

	app.New(cfg)

	// app.New(log, cfg.APP.Port)

	// e.Logger.Fatal(e.Start(":1323"))
}
