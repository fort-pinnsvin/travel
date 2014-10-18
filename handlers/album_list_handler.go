package handlers

import (
	"github.com/martini-contrib/sessions"
	"github.com/martini-contrib/render"
	"github.com/go-martini/martini"
	"github.com/fort-pinnsvin/travel/helpfunc"

	"github.com/fort-pinnsvin/travel/models"
	"labix.org/v2/mgo/bson"
	"fmt"
)

func AlbumListHandler(rnd render.Render, session sessions.Session,  params martini.Params){
	if session.Get("auth_id") != "" {
		user_auth := helpfunc.GetAuthUser(session)
		albums := []models.Marker{}
		query := make(bson.M)
		query["owner"] = user_auth.Id
		models.MarkerCollection.Find(query).Limit(1024).Iter().All(&albums)
		fmt.Println(albums)
		rnd.HTML(200, "lists_album", map[string]interface {}{
				"user" : user_auth,
				"albums" : albums,
		});
	}
}
