package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"academy/internal/app"
)

func bodyDumpConfig() middleware.BodyDumpConfig {
	handler := func(c echo.Context, req []byte, res []byte) {
		log.Println("headers:", c.Request().Header)
		log.Println("request:", string(req))
		log.Println("response:", string(res))
	}
	return middleware.BodyDumpConfig{Handler: handler}
}

func notFoundOnProduction(cv *app.Configs) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if cv.State != "dev" {
				return c.NoContent(http.StatusNotFound)
			}
			return next(c)
		}
	}
}

func moviesEndpointEnable(cv *app.Configs) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if !cv.Movie.Enable {
				log.Println("movie endpoint disabled")
				return c.NoContent(http.StatusNotFound)
			}
			return next(c)
		}
	}
}
