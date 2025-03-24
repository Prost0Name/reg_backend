package main

import (
	"backend/internal/app"
	"backend/internal/config"
	"backend/internal/model"
	"fmt"
)

const (
	envLocal = "local"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg.JwtSecret)

	model.InitDatabase(cfg.DSN)

	app.New(cfg)

	// app.New(log, cfg.APP.Port)

	// e.Logger.Fatal(e.Start(":1323"))
}
