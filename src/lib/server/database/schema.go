package database

import "time"

type Movie struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Year        string    `json:"year"`
	Directors   []string  `json:"directors"`
	Writers     []string  `json:"writers"`
	Producers   []string  `json:"producers"`
	Editors     []string  `json:"editors"`
	Cameras     []string  `json:"cameras"`
	Genres      []string  `json:"genres"`
	ReleaseDate time.Time `json:"release_date"`
	Countries   []string  `json:"countries"`
	Duration    float64   `json:"duration"`
	Description string    `json:"description"`
	Poster      string    `json:"poster"`
	Images      []string  `json:"images"`
	Actors      []string  `json:"actors"`
	Created     time.Time `json:"created"`
}

type Series struct {
}

type Actor struct {
	Id         string    `bson:"id"`
	Name       string    `bson:"name"`
	Surname    string    `bson:"surname"`
	Age        int       `bson:"age"`
	Birthday   time.Time `bson:"birthday"`
	Gender     string    `bson:"gender"`
	Birthplace string    `bson:"birthplace"`
	Bio        string    `bson:"bio"`
	Images     []string  `bson:"images"`
	Created    time.Time `bson:"created"`
}

type Article struct {
	Id      string    `bson:"id"`
	Title   string    `bson:"title"`
	Content string    `bson:"content"`
	Author  string    `bson:"author"`
	Image   string    `bson:"image"`
	Created time.Time `bson:"created"`
}
