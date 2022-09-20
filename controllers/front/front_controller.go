package front

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// TODO
func ContainerPage(c echo.Context) error {
	return c.Render(http.StatusOK, "containers.html", map[string]interface{}{
		"pageName": "Container",
	})
}

// TODO
func VolumePage(c echo.Context) error {
	return c.Render(http.StatusOK, "volumes.html", map[string]interface{}{
		"pageName": "Volume",
	})
}

// TODO
func ImagePage(c echo.Context) error {
	return c.Render(http.StatusOK, "images.html", map[string]interface{}{
		"pageName": "Image",
	})
}
