package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Response struct {
	Success bool        `json:"success"`
	Code    int         `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type Movie struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Title    string             `bson:"title,omitempty"`
	Plot     string             `bson:"plot,omitempty"`
	FullPlot string             `bson:"fullplot,omitempty"`
	Genres   []string           `bson:"genres,omitempty"`
	Cast     []string           `bson:"cast,omitempty"`
	Rated    string             `bson:"rated,omitempty"`
	Year     int                `bson:"year,omitempty"`
	Award    Award              `bson:"award,omitempty"`
}

type Award struct {
	Wins        int    `bson:"wins"`
	Nominations int    `bson:"nominations"`
	Text        string `bson:"text"`
}

type Comment struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	MovieID primitive.ObjectID `bson:"movie_id,omitempty"`
	Name    string             `bson:"name,omitempty"`
	Email   string             `bson:"email,omitempty"`
	Text    string             `bson:"text,omitempty"`
	Date    time.Time          `bson:"date,omitempty"`
}
