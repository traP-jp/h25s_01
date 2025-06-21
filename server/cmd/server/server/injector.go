package server

import (
	"backend/internal/handler"
	"backend/internal/infrastructure/database"
	"backend/internal/infrastructure/file"
	"backend/internal/router"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type Server struct {
	Router *echo.Echo
}

func Inject(db *sqlx.DB) *Server {
	shopRepo := database.NewShopRepository(db)
	reviewRepo := database.NewReviewRepository(db)
	stationRepo := database.NewStationRepository(db)
	fileRepo, err := file.NewFileRepository()
	if err != nil {
		panic("failed to create file repository: " + err.Error())
	}

	shopHandler := handler.NewShopHandler(shopRepo, fileRepo)
	reviewHandler := handler.NewReviewHandler(reviewRepo, fileRepo)
	stationHandler := handler.NewStationHandler(stationRepo, shopRepo)
	fileHandler := handler.NewFileHandler(fileRepo)

	echoRouter := router.NewRouter(
		shopHandler,
		reviewHandler,
		stationHandler,
		fileHandler,
	)

	return &Server{
		Router: echoRouter,
	}
}
