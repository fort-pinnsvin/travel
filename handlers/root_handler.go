package handlers

import (
	"github.com/fort-pinnsvin/travel/models"
	"github.com/martini-contrib/oauth2"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"net/http"
)

func Root(tokens oauth2.Tokens, rnd render.Render, r *http.Request, session sessions.Session) {
	if tokens.IsExpired() {
		rnd.HTML(200, "home", nil)
	} else {
		user := &models.User{}
		user.FirstName = session.Get("first_name").(string)
		user.LastName = session.Get("last_name").(string)
		user.Id = session.Get("auth_id").(string)
		user.Avatar = session.Get("avatar").(string)
		rnd.HTML(200, "home_user", user)
	}
}
