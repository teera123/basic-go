package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"academy/internal/app"
	"academy/internal/movies"
	"academy/internal/ping"
)

type route struct {
	Group          string
	Path           string
	HttpMethod     string
	HandlerFunc    echo.HandlerFunc
	MiddlewareFunc []echo.MiddlewareFunc
}

func NewRoutes(e *echo.Echo, cv *app.Configs) error {
	movieRepo := movies.NewMongoRepo(cv)
	movieSrv := movies.NewService(cv, movieRepo)
	movieEndpoint := movies.NewEndpoint(cv, movieSrv)

	routes := []route{
		{
			Group:          "",
			Path:           "/health_check",
			HttpMethod:     http.MethodGet,
			HandlerFunc:    ping.HealthCheck,
			MiddlewareFunc: nil,
		},
		{
			Group:          "movies",
			Path:           "/year",
			HttpMethod:     http.MethodPost,
			HandlerFunc:    movieEndpoint.GetMoviesByYear,
			MiddlewareFunc: []echo.MiddlewareFunc{moviesEndpointEnable(cv)},
		},
	}

	// config middleware
	e.Use(middleware.Recover())
	e.Use(middleware.BodyDumpWithConfig(bodyDumpConfig()))

	// http connection
	for _, r := range routes {
		e.Group(r.Group).Add(r.HttpMethod, r.Path, r.HandlerFunc, r.MiddlewareFunc...)
	}

	return nil
}
