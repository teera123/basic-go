package movies

import (
	"context"

	"github.com/labstack/echo"

	"academy/internal/app"
	"academy/internal/utils"
)

type Endpoint struct {
	cv       *app.Configs
	movieSrv movieService
}

func NewEndpoint(cv *app.Configs, movieSrv movieService) *Endpoint {
	return &Endpoint{cv: cv, movieSrv: movieSrv}
}

type movieService interface {
	GetMoviesByYear(ctx context.Context, start, end, limit int) (GetMoviesByYearResponse, error)
}

func (e Endpoint) GetMoviesByYear(c echo.Context) error {
	var req struct {
		Start int `json:"start"`
		End   int `json:"end"`
		Limit int `json:"limit"`
	}
	if err := c.Bind(&req); err != nil {
		return utils.ResponseJSON(c, nil, err)
	}
	res, err := e.movieSrv.GetMoviesByYear(c.Request().Context(), req.Start, req.End, req.Limit)
	return utils.ResponseJSON(c, res, err)
}
