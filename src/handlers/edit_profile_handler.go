package handlers

import (
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"models"
)

func Edit(rnd render.Render, session sessions.Session) {
	if session.Get("auth_id") != "" {
		userData := &models.User{}
		models.UserCollection.FindId(session.Get("auth_id")).One(&userData)

		rnd.HTML(200, "edit_profile", userData)
	}
}
