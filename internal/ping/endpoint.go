package ping

import (
	"net/http"

	"github.com/labstack/echo"
)

func HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"status": "OK"})
}
