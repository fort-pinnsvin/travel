package models

import (
	"crypto/rand"
	"fmt"
	"labix.org/v2/mgo"
	"os"
)

var UserCollection *mgo.Collection

func GenerateId() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
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
}
