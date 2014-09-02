package main

import (
	"github.com/go-martini/martini"
	"travel/handlers"
	"github.com/martini-contrib/render"
	"travel/models"
)

func main() {

	models.CreadeDB()

	m := martini.Classic()

	m.Use(martini.Static("assets"))

	m.Use(render.Renderer(render.Options{
		Directory:  "templates",                         // Specify what path to load the templates from.
		//Layout:     "layout",                            // Specify a layout template. Layouts can call {{ yield }} to render the current template.
		Extensions: []string{".tmpl", ".html"},          // Specify extensions to load for templates.
		//Funcs:      []template.FuncMap{unescapeFuncMap}, // Specify helper function maps for templates to access.
		Charset:    "UTF-8",                             // Sets encoding for json and html content-types. Default is "UTF-8".
		IndentJSON: true,                                // Output human readable JSON
}))

	m.Get("/", handlers.MainHandler)
	m.Get("/regestration", handlers.RegestrationHandler)
	m.Post("/signup", handlers.SignupHandler)
	m.Run()
}
