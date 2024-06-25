package database

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Movie struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Year        string             `bson:"year"`
	Directors   []string           `bson:"directors"`
	Writers     []string           `bson:"writers"`
	Producers   []string           `bson:"producers"`
	Editors     []string           `bson:"editors"`
	Cameras     []string           `bson:"cameras"`
	Genres      []string           `bson:"genres"`
	ReleaseDate time.Time          `bson:"release_date"`
	Countries   []string           `bson:"countries"`
	Duration    float64            `bson:"duration"`
	Description string             `bson:"description"`
	Poster      string             `bson:"poster"`
	Images      []string           `bson:"images"`
	Actors      []string           `bson:"actors"`
	Created     time.Time          `bson:"created"`
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
