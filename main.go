package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"handlers"
	"models"
	"github.com/martini-contrib/sessions"
)

func main() {

	models.CreadeDB()

	store := sessions.NewCookieStore([]byte("secret123"))

	m := martini.Classic()

	m.Use(martini.Static("assets"))

	m.Use(sessions.Sessions("auth", store))

	m.Use(render.Renderer(render.Options{
		Directory: "templates", // Specify what path to load the templates from.
		//Layout:     "layout",                            // Specify a layout template. Layouts can call {{ yield }} to render the current template.
		Extensions: []string{".tmpl", ".html"}, // Specify extensions to load for templates.
		//Funcs:      []template.FuncMap{unescapeFuncMap}, // Specify helper function maps for templates to access.
		Charset:    "UTF-8", // Sets encoding for json and html content-types. Default is "UTF-8".
		IndentJSON: true,    // Output human readable JSON
	}))

	m.Get("/", handlers.MainHandler)
	m.Get("/registration", handlers.RegestrationHandler)
	m.Post("/signup", handlers.SignupHandler)
	m.Post("/signin", handlers.SigninHandler)
	m.Get("/login", handlers.LoginHandler)
	m.Get("/user/:id", handlers.UserHandler)

	m.Run()
}
