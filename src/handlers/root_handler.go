package handlers

import (
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"net/http"
	"github.com/martini-contrib/oauth2"
	"models"
)

func Root(tokens oauth2.Tokens, rnd render.Render, r *http.Request, session sessions.Session) {
	user := &models.User{}
	user.FirstName = session.Get("first_name").(string)
	user.LastName = session.Get("last_name").(string)
	user.Id = session.Get("auth_id").(string)
	if tokens.IsExpired(){
		rnd.HTML(200, "home", user)
	}else{
		rnd.HTML(200, "home_user", user)
	}
}
