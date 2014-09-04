package test

import (
	"fmt"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"models"
	"net/http"
)

func Root(rnd render.Render, r *http.Request, session sessions.Session) {
	rnd.HTML(200, "home", nil)
}
