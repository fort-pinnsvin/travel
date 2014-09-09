package models

import (
	"labix.org/v2/mgo"
	"os"
)

type User struct {
	Id         string `bson:"_id,omitempty"`
	FirstName  string
	LastName   string
	Email      string
	Avatar     string
	Birthday   string
	Country    string
	Status     string
	About      string
}

type Marker struct {
	Id        string `bson:"_id,omitempty"`
	Owner     string
	Name      string
	Latitude  string
	Longitude string
}

type Post struct {
	Id        string `bson:"_id,omitempty"`
	Owner     string
	Title     string
	Text      string
	Date      string
	Nano      int
	Like      int
	OwnerUser User
}

const Layout = "Jan 2, 2006 at 3:04pm"

type ByPost []Post

func (a ByPost) Len() int { return len(a) }
func (a ByPost) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByPost) Less(i, j int) bool { return a[i].Nano < a[j].Nano }

type FollowEdge struct {
	Id        string `bson:"_id,omitempty"`
	Follower  string
	Following string
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
	PostCollection = session.DB(database).C("post")
	FollowCollection = session.DB(database).C("followers")
}
