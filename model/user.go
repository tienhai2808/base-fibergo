package model

import "time"

type User struct {
	Username  string    `json:"username" bson:"username"`
	Password  string    `json:"-" bson:"password"`
	Email     string    `json:"email" bson:"email"`
	FirstName string    `json:"first_name" bson:"first_name"`
	LastName  string    `json:"last_name" bson:"last_name"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"-" bson:"updated_at"`
}
