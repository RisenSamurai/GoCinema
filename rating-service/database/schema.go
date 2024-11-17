package database

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type DetailedMovie struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Directors   []string  `bson:"directors" json:"directors"`
	Writers     []string  `bson:"writers" json:"writers"`
	Producers   []string  `bson:"producers" json:"producers"`
	Editors     []string  `bson:"editors" json:"editors"`
	Cameras     []string  `bson:"cameras" json:"cameras"`
	Genres      []string  `bson:"genres" json:"genres"`
	Actors      []string  `bson:"actors" json:"actors"`
	Country     []string  `json:"origin_country"`
	Keywords    []string  `bson:"keywords" json:"keywords"`
	Budget      float64   `bson:"budget" json:"budget"`
	Language    string    `json:"original_language"`
	ReleaseDate time.Time `json:"release_date"`
	Duration    float64   `bson:"duration" json:"duration"`
	Description string    `json:"overview"`
	Popularity  float64   `json:"popularity"`
	VoteAverage float64   `json:"vote_average"`
	VoteCount   float64   `json:"vote_count"`
	Revenue     float64   `json:"revenue"`
	Status      string    `json:"status"`
	Tagline     string    `json:"tagline"`
	Poster      string    `json:"poster_path"`
	Images      []string  `bson:"images" json:"backdrop_path"`
	Created     time.Time `bson:"created" json:"created"`
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
