package models

type UserDocument struct {
	Id 			string `bson:"_id,omitempty"`
	FirstName 	string
	LastName	string
	Email		string
	Password 	string
}
