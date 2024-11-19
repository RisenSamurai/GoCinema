package database

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"rating_microservice/util"
	"time"
)

type DetailedMovie struct {
	Id                  int          `json:"id"`
	Title               string       `json:"title"`
	Country             []any        `json:"production_countries"`
	Budget              float64      `json:"budget"`
	Language            string       `json:"original_language"`
	ReleaseDate         string       `json:"release_date"`
	Duration            float64      `json:"runtime"`
	Description         string       `json:"overview"`
	Popularity          float64      `json:"popularity"`
	VoteAverage         float64      `json:"vote_average"`
	VoteCount           float64      `json:"vote_count"`
	Revenue             float64      `json:"revenue"`
	Status              string       `json:"status"`
	Tagline             string       `json:"tagline"`
	Poster              string       `json:"poster_path"`
	Genres              []util.Genre `json:"genres"`
	ProductionCompanies []any        `json:"production_companies"`
}

type MainPageMovie struct {
	Id         int    `json:"id"`
	PosterPath string `json:"poster_path"`
	Title      string `json:"title"`
}

type Actor struct {
	Id         primitive.ObjectID `bson:"_id,omitempty"`
	Name       string             `bson:"name"`
	Surname    string             `bson:"surname"`
	Age        int                `bson:"age"`
	Birthday   time.Time          `bson:"birthday"`
	Gender     string             `bson:"gender"`
	Birthplace string             `bson:"birthplace"`
	Bio        string             `bson:"bio"`
	Images     []string           `bson:"images"`
	Created    time.Time          `bson:"created"`
}

type Article struct {
	Id      primitive.ObjectID `bson:"_id,omitempty"`
	Title   string             `bson:"title"`
	Content string             `bson:"content"`
	Author  string             `bson:"author"`
	Image   string             `bson:"image"`
	Created time.Time          `bson:"created"`
}
