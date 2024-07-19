package database

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Movie struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name" json:"name"`
	Year        string             `bson:"year" json:"year"`
	Directors   []string           `bson:"directors" json:"directors"`
	Writers     []string           `bson:"writers" json:"writers"`
	Producers   []string           `bson:"producers" json:"producers"`
	Editors     []string           `bson:"editors" json:"editors"`
	Cameras     []string           `bson:"cameras" json:"cameras"`
	Genres      []string           `bson:"genres" json:"genres"`
	Actors      []string           `bson:"actors" json:"actors"`
	Countries   []string           `bson:"countries" json:"countries"`
	Keywords    []string           `bson:"keywords" json:"keywords"`
	Budget      float64            `bson:"budget" json:"budget"`
	Language    string             `bson:"language" json:"language"`
	ReleaseDate time.Time          `bson:"release_date" json:"release_date"`
	Duration    float64            `bson:"duration" json:"duration"`
	Description string             `bson:"description" json:"description"`
	Poster      string             `bson:"poster" json:"poster"`
	Images      []string           `bson:"images" json:"images"`
	Created     time.Time          `bson:"created" json:"created"`
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
