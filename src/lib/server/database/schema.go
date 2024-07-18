package database

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Movie struct {
	Name        string    `json:"name"`
	Year        string    `json:"year"`
	Directors   []string  `json:"directors"`
	Writers     []string  `json:"writers"`
	Producers   []string  `json:"producers"`
	Editors     []string  `json:"editors"`
	Cameras     []string  `json:"cameras"`
	Genres      []string  `json:"genres"`
	Actors      []string  `json:"actors"`
	Countries   []string  `json:"countries"`
	Keywords    []string  `json:"keywords"`
	Budget      float64   `json:"budget"`
	Language    string    `json:"language"`
	ReleaseDate time.Time `json:"release_date"`
	Duration    float64   `json:"duration"`
	Description string    `json:"description"`
	Poster      string    `json:"poster"`
	Images      []string  `json:"images"`
	Created     time.Time `json:"created"`
}

type Series struct {
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
