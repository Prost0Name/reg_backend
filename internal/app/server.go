package app

import (
	"backend/internal/app/middleware"
	"backend/internal/app/routes"
	"log"

	"github.com/labstack/echo/v4"
)

func New(port string) {
	e := echo.New()

	middleware.CORS(e)

	routes.Users(e)
	// middleware.New(e)

	if err := e.Start(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
