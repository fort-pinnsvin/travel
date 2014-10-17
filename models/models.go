package models

import (
	"labix.org/v2/mgo"
	"github.com/fort-pinnsvin/travel/utils"
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
	Language   string
}

type Marker struct {
	Id        string `bson:"_id,omitempty"`
	Owner     string
	Name      string
	Latitude  string
	Longitude string
	Description string
}

type Post struct {
	Id        string `bson:"_id,omitempty"`
	Owner     string
	Title     string
	Text      string
	Date      string
	Nano      int64
	Like      int
	OwnerUser User
	IsLiked bool
}

type Like struct {
	Liker        string
	IdPost       string

}

const Layout = "Jan 2, 2006 at 3:04pm"

type ByPost []Post

func (a ByPost) Len() int { return len(a) }
func (a ByPost) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByPost) Less(i, j int) bool { return a[i].Nano > a[j].Nano }

type FollowEdge struct {
	Id        string `bson:"_id,omitempty"`
	Follower  string
	Following string
}

func ConnectToDataBase() {
	url := utils.GetValue("DB_URL", "localhost")
	database := utils.GetValue("DB", "travel")

	session, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}
	UserCollection = session.DB(database).C("users")
	MarkerCollection = session.DB(database).C("markers")
	PostCollection = session.DB(database).C("post")
	FollowCollection = session.DB(database).C("followers")
	LikeCollection = session.DB(database).C("like")
}
