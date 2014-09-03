package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"handlers"
	"models"
)

func main() {
	models.ConnectToDataBase()
	store := sessions.NewCookieStore([]byte("yLiCQYG7CAflDavqGH461IO0MHp7TEbpg6TwHBWdJzNwYod1i5ZTbrIF5bEoO3oP"))
	m := martini.Classic()
	m.Use(martini.Static("assets"))
	m.Use(sessions.Sessions("auth", store))
	m.Use(render.Renderer(render.Options{
		Directory:  "templates",
		Extensions: []string{".tmpl", ".html"},
		Charset:    "UTF-8",
		IndentJSON: true,
	}))

	m.Get("/", handlers.Root)
	m.Get("/signup", handlers.SignUpForm)
	m.Post("/signup", handlers.PostSignUp)
	m.Get("/signin", handlers.SignInForm)
	m.Post("/signin", handlers.PostSignIn)
	m.Get("/user/:id", handlers.UserProfile)
	m.Get("/logout", handlers.Logout)

	m.Run()
}
