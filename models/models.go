package models

import (
	"github.com/fort-pinnsvin/travel/utils"
	"html/template"
	"labix.org/v2/mgo"
)

type User struct {
	Id        string `bson:"_id,omitempty"`
	FirstName string
	LastName  string
	Email     string
	Avatar    string
	Birthday  string
	Country   string
	Status    string
	About     string
	Language  string
	Latitude  float64
	Longitude float64
	Points 	  float64
}

type Marker struct {
	Id          string `bson:"_id,omitempty"`
	Owner       string
	Name        string
	Latitude    string
	Longitude   string
	Description string
	FullAddress string
	Date        string
	Nano        int64
	Country		string
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
	IsLiked   bool
	Html      template.HTML
}

type Like struct {
	Liker  string
	IdPost string
}

type Photo struct {
	AlbumId string
	Name    string
}

type Country struct {
	Code  string `bson:"_id,omitempty"`
	Count int
	Name  string
}

type PostBlog struct {
	Id     string `bson:"_id,omitempty"`
	IdBlog string
	Owner  string
	Title  string
	Text   string
	Date   string
	Nano   int64
}

type Blog struct {
	Id    string `bson:"_id,omitempty"`
	Owner string
	Name  string
	Date  string
	Nano  int64
}

const Layout = "Jan 2, 2006 at 3:04pm"

// Sort Posts by Time
type ByPost []Post

func (a ByPost) Len() int           { return len(a) }
func (a ByPost) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPost) Less(i, j int) bool { return a[i].Nano > a[j].Nano }

type ByPostBlog []PostBlog

func (a ByPostBlog) Len() int           { return len(a) }
func (a ByPostBlog) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPostBlog) Less(i, j int) bool { return a[i].Nano < a[j].Nano }

type ByBlog []Blog

func (a ByBlog) Len() int           { return len(a) }
func (a ByBlog) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByBlog) Less(i, j int) bool { return a[i].Nano > a[j].Nano }

// Sort Countryes by Count
type ByCountry []Country

func (a ByCountry) Len() int           { return len(a) }
func (a ByCountry) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCountry) Less(i, j int) bool { return a[i].Count > a[j].Count }

// Sort Countryes by Count
type ByUser []User

func (a ByUser) Len() int           { return len(a) }
func (a ByUser) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByUser) Less(i, j int) bool { return a[i].Points > a[j].Points }

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
	PhotoCollection = session.DB(database).C("photo")
	CountryCollection = session.DB(database).C("country")
	PostBlogCollection = session.DB(database).C("blog_post")
	BlogCollection = session.DB(database).C("blog")
}
