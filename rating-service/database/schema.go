package database

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

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
