package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"strconv"
)

const UserIDKey = "userId"

func UserIDMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.Request().Header.Get("X-Forwarded-User")
		if username != "" {
			c.Set(UserIDKey, username)

			return next(c)
		}

		isDebug, err := strconv.ParseBool(os.Getenv("DEBUG"))
		if err != nil || !isDebug {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}

		c.Set(UserIDKey, "traP")

		return next(c)
	}
}
