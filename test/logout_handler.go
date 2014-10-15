package test

import (
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"net/http"
)

func Logout(rnd render.Render, r *http.Request, session sessions.Session) {
	session.Clear()
	rnd.Redirect("/")
}
