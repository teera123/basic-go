package movies

import (
	"context"

	"academy/internal/app"
	"academy/internal/model"
)

type Service struct {
	cv        *app.Configs
	mongoRepo mongoRepo
}

type mongoRepo interface {
	GetMoviesByYear(ctx context.Context, start, end, limit int) ([]model.Movie, error)
}

func NewService(cv *app.Configs, repo mongoRepo) *Service {
	return &Service{cv: cv, mongoRepo: repo}
}

type GetMoviesByYearResponse struct {
	Length int           `json:"length"`
	Data   []model.Movie `json:"data"`
}

func (s Service) GetMoviesByYear(ctx context.Context, start, end, limit int) (GetMoviesByYearResponse, error) {
	data, err := s.mongoRepo.GetMoviesByYear(ctx, start, end, limit)
	if err != nil {
		return GetMoviesByYearResponse{}, err
	}
	return GetMoviesByYearResponse{
		Length: len(data),
		Data:   data,
	}, nil
}
