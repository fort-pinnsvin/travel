package handlers

import (
	"github.com/martini-contrib/sessions"
	"github.com/martini-contrib/render"
	"github.com/go-martini/martini"
	"github.com/fort-pinnsvin/travel/helpfunc"

)

func AlbumListHandler(rnd render.Render, session sessions.Session,  params martini.Params){
	if session.Get("auth_id") != "" {
		user := helpfunc.GetAuthUser(session)
		rnd.HTML(200, "album_list", user);
	}
}
