package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetUserID(c echo.Context) (string, error) {
	userID, ok := c.Get(UserIDKey).(string)
	if !ok {
		return "", echo.NewHTTPError(http.StatusUnauthorized, "user ID not found in context")
	}

	return userID, nil
}

func errorResponse(c echo.Context, status int, msg string) error {
	return c.JSON(status, map[string]string{"error": msg})
}
