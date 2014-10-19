package handlers

import (
	"github.com/fort-pinnsvin/travel/helpfunc"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"fmt"
	"github.com/fort-pinnsvin/travel/models"
	"labix.org/v2/mgo/bson"
)

func AlbumListHandler(rnd render.Render, session sessions.Session, params martini.Params) {
	if session.Get("auth_id") != "" {
		user_auth := helpfunc.GetAuthUser(session)
		albums := []models.Marker{}
		query := make(bson.M)
		query["owner"] = user_auth.Id
		models.MarkerCollection.Find(query).Limit(1024).Iter().All(&albums)
		for i := 0; i < len(albums); i ++ {
			if len(albums[i].Name) > 15 {
				albums[i].Name = albums[i].Name[0:15] + "..."
			}
		}
		fmt.Println(albums)
		rnd.HTML(200, "lists_album", map[string]interface{}{
			"user":   user_auth,
			"albums": albums,
		})
	}
}
