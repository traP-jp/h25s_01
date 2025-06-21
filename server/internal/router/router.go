package router

import (
	"net/http"

	"backend/internal/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(
	shopHandler *handler.ShopHandler,
	reviewHandler *handler.ReviewHandler,
	stationHandler *handler.StationHandler,
) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.Use(handler.UserIDMiddleware)

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "ok",
		})
	})

	api := e.Group("/api/v1")
	{
		shops := api.Group("/shops")
		{
			shops.GET("", shopHandler.GetShops)
			shops.POST("", shopHandler.CreateShop)
			shops.GET("/:id", shopHandler.GetShopDetail)
			shops.PUT("/:id", shopHandler.UpdateShop)
			shops.DELETE("/:id", shopHandler.Delete)
			shops.POST("/:id/images", shopHandler.ShopImgUpload)
			shops.DELETE("/:id/images", shopHandler.DeletePicture)
		}

		reviews := api.Group("/reviews")
		{
			reviews.GET("", reviewHandler.GetReviews)
			reviews.POST("", reviewHandler.CreateReview)
			reviews.GET("/:id", reviewHandler.GetReview)
			reviews.PUT("/:id", reviewHandler.UpdateReview)
			reviews.DELETE("/:id", reviewHandler.DeleteReview)
			reviews.POST("/:id/images", reviewHandler.UploadImage)
		}

		stations := api.Group("/stations")
		{
			stations.GET("", stationHandler.GetStations)
			stations.POST("", stationHandler.CreateStation)
			stations.GET("/:id", stationHandler.GetStationDetail)
			stations.PUT("/:id", stationHandler.UpdateStation)
			stations.DELETE("/:id", stationHandler.DeleteStation)
			stations.GET("/:id/shops", stationHandler.GetShopAroundStation)
		}

	}

	return e
}
