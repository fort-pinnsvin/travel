package  models

import (
	"labix.org/v2/mgo"
	"crypto/rand"
	"fmt"
)
var UserCollection *mgo.Collection

func GenerateId() string{
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x",b)
}

func CreadeDB(){
	session, err := mgo.Dial("localhost")

	if err != nil{
		panic(err)
	}

	UserCollection = session.DB("travel").C("users")
}
