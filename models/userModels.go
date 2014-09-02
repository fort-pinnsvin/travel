package models

type UserDocument struct {
	Id 			string `bson:"_id,omitempty"`
	FirstName 	string
	LastName	string
	Emain		string
	Password 	string
}
