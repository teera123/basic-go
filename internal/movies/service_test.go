package movies

import (
	"context"
	"testing"

	"academy/internal/model"
)

var fakeDB = []model.Movie{
	{Title: "test movie 1", Year: 1991},
	{Title: "test movie 2", Year: 2000},
	{Title: "test movie 3", Year: 2001},
}

type fakeMongo struct{}

func (f fakeMongo) GetMoviesByYear(ctx context.Context, start, end, limit int) ([]model.Movie, error) {
	var years []int
	for i := start; i <= end; i++ {
		years = append(years, i)
	}

	var rtn []model.Movie
	for _, y := range years {
		for _, f := range fakeDB {
			if f.Year == y {
				rtn = append(rtn, f)
			}
		}
	}
	return rtn, nil
}

func TestService_GetMoviesByYear(t *testing.T) {
	tcs := []struct {
		start       int
		end         int
		expectedLen int
	}{
		{
			start:       1991,
			end:         1992,
			expectedLen: 1,
		},
		{
			start:       1990,
			end:         2020,
			expectedLen: 3,
		},
	}

	srv := NewService(nil, fakeMongo{})
	for _, tc := range tcs {
		data, err := srv.GetMoviesByYear(nil, tc.start, tc.end, 10)
		if err != nil {
			t.Error("found error")
			continue
		}
		if data.Length != tc.expectedLen {
			t.Errorf("expected length %d, but got %d", tc.expectedLen, data.Length)
		}
	}
}
