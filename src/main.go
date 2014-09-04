package main

import (
	"github.com/go-martini/martini"
	gooauth2 "github.com/golang/oauth2"
	"github.com/martini-contrib/oauth2"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"models"
	"handlers"
)

func main() {
	models.ConnectToDataBase()

	store := sessions.NewCookieStore([]byte("yLiCQYG7CAflDavqGH461IO0MHp7TEbpg6TwHBWdJzNwYod1i5ZTbrIF5bEoO3oP"))
	m := martini.Classic()
	m.Use(martini.Static("assets"))
	m.Use(sessions.Sessions("auth", store))
	m.Use(oauth2.Google(&gooauth2.Options{
		ClientID:     "903364406910-m1b4j2vjkfd3qj1npusv6p2qk38fqb3q",
		ClientSecret: "iofaFDfJuJRkPTjPu4NuHx61",
		RedirectURL:  "http://localhost:3000/oauth2callback",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile"},
	}))
	m.Use(render.Renderer(render.Options{
		Directory:  "templates",
		Extensions: []string{".tmpl", ".html"},
		Charset:    "UTF-8",
		IndentJSON: true,
	}))

	m.Get("/", handlers.Root)
	m.Get("/data", oauth2.LoginRequired, handlers.GetData)

	// Routes that require a logged in user
	// can be protected with oauth2.LoginRequired handler.
	// If the user is not authenticated, they will be
	// redirected to the login path.

	m.Run()
}
