package front

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// TODO
func HomePage(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", "World")
}