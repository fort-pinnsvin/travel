package test

import (
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"net/http"
)

func Logout(rnd render.Render, r *http.Request, session sessions.Session) {
	if session.Get("auth") != "" && session.Get("auth_id") != "" {
		session.Set("auth", "")
		session.Set("auth_id", "")
		rnd.Redirect("/")
	}
}
