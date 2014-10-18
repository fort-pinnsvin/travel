package models

import (
	"crypto/rand"
	"fmt"
	"labix.org/v2/mgo"
)

var UserCollection 		*mgo.Collection
var MarkerCollection 	*mgo.Collection
var PostCollection 		*mgo.Collection
var LikeCollection 		*mgo.Collection
var FollowCollection 	*mgo.Collection
var PhotoCollection 	*mgo.Collection
var PostBlogCollection	*mgo.Collection
var BlogCollection		*mgo.Collection

func GenerateId() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
