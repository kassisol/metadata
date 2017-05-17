package api

import (
	"net/http"

	"github.com/labstack/echo"
)

func indexHandle(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}
