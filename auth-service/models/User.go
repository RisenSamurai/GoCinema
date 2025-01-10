package models

type User struct {
	id       string `bson:"_id,omitempty" json:"id"`
	email    string `bson:"email" json:"email"`
	password string `bson:"password" json:"password"`
	role     string `bson:"role" json:"role"`
}
