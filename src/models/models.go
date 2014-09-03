package models

type User struct {
	Id        string `bson:"_id,omitempty"`
	FirstName string
	LastName  string
	Email     string
	Password  string
}
