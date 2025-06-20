package server

import (
	"backend/internal/handler"
	"backend/internal/infrastructure/database"
	"backend/internal/router"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type Server struct {
	Router *echo.Echo
}

func Inject(db *sqlx.DB) *Server {
	userRepo := database.NewUserRepository(db)

	userHandler := handler.NewUserHandler(userRepo)

	echoRouter := router.NewRouter(userHandler)

	return &Server{
		Router: echoRouter,
	}
}
