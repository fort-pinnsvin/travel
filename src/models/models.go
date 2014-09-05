package models

import (
	"labix.org/v2/mgo"
	"os"
)

type User struct {
	Id        string `bson:"_id,omitempty"`
	FirstName string
	LastName  string
	Email     string
	Avatar    string
}

type Marker struct {
	Id        string `bson:"_id,omitempty"`
	Owner     string
	Name      string
	Latitude  string
	Longitude string
}

func ConnectToDataBase() {
	url := os.Getenv("DB_URL")
	if url == "" {
		url = "localhost"
	}
	database := os.Getenv("DB")
	if database == "" {
		database = "travel"
	}

	session, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}
	UserCollection = session.DB(database).C("users")
	MarkerCollection = session.DB(database).C("markers")
}
