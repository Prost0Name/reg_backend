package main

import (
	"backend/internal/models"
	"backend/internal/middleware"
	"backend/internal/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	models.InitDatabase()
	middleware.New(e)
	routes.New(e)

	e.Logger.Fatal(e.Start(":1323"))
}
