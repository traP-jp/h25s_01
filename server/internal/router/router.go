package router

import (
	"net/http"

	"backend/internal/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(userHandler *handler.UserHandler) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.Use(userIDMiddleware)

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "ok",
		})
	})

	api := e.Group("/api/v1")
	{
		users := api.Group("/users")
		{
			users.POST("", userHandler.CreateUser)
			users.GET("/:id", userHandler.GetUser)
			users.DELETE("/:id", userHandler.DeleteUser)
		}
	}

	return e
}

func GetUserID(c echo.Context) (string, error) {
	userID, ok := c.Get(userIDKey).(string)
	if !ok {
		return "", echo.NewHTTPError(http.StatusUnauthorized, "user ID not found in context")
	}

	return userID, nil
}
