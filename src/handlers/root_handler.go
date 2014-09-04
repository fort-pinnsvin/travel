package handlers

import (
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"net/http"
	"github.com/martini-contrib/oauth2"
)

func Root(tokens oauth2.Tokens, rnd render.Render, r *http.Request, session sessions.Session) {
	s := ""
	if tokens.IsExpired(){
		s = "Go to /login plase, you are not signin."
	}else{
		s = "You are signin"
	}
	rnd.HTML(200, "home", s)
}
