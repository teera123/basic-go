package movies

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"academy/internal/app"
	"academy/internal/model"
)

const mflixDB = "sample_mflix"

type repo struct {
	cv *app.Configs
}

func NewMongoRepo(cv *app.Configs) *repo {
	return &repo{cv: cv}
}

func (r repo) GetMoviesByYear(ctx context.Context, start, end, limit int) ([]model.Movie, error) {
	if limit > 10 || limit == 0 {
		limit = 10
	}

	coll := r.cv.MongoDB.Client.Database(mflixDB).Collection("movies")
	cur, err := coll.Find(ctx, bson.M{"year": bson.M{"$gt": start, "$lt": end}}, options.Find().SetLimit(int64(limit)))
	if err != nil {
		return nil, err
	}

	var rtn []model.Movie
	if err := cur.All(ctx, &rtn); err != nil {
		return nil, err
	}
	return rtn, nil
}
