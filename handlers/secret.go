package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

func GetSecret(c echo.Context) error {
	return c.String(http.StatusOK, "SUCCESS")
}
