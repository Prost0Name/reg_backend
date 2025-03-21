package app

import (
	"backend/internal/app/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New(port string) {
	e := echo.New()

	e.Use(middleware.Logger())

	routes.Users(e)
	// middleware.New(e)
	// middleware.cors(e)

	if err := e.Start(":" + port); err != nil {
		e.Logger.Fatal("Failed to start server: ", err)
	}
}
