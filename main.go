package main

import (
	"github.com/go-martini/martini"
	gooauth2 "github.com/golang/oauth2"
	"github.com/martini-contrib/oauth2"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"Travel/handlers"
	"Travel/models"
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
	m.Get("/user/:id", oauth2.LoginRequired, handlers.UserProfile)
	m.Get("/signin", oauth2.LoginRequired, handlers.GetData)
	m.Get("/edit", oauth2.LoginRequired, handlers.Edit)
	m.Post("/update", oauth2.LoginRequired, handlers.EditPost)
	m.Get("/search", oauth2.LoginRequired, handlers.Search)
	m.Get("/markers", oauth2.LoginRequired, handlers.GetMarkers)
	m.Get("/markers/create", oauth2.LoginRequired, handlers.CreateMarker)
	m.Get("/css/bootstrap.min.css", handlers.LoadTheme)
	m.Post("/set_theme", handlers.SetTheme)
	m.Post("/save_post",oauth2.LoginRequired, handlers.SavePost)
	m.Post("/remove_post",oauth2.LoginRequired, handlers.RemovePost)
	m.Post("/add_like",oauth2.LoginRequired, handlers.AddLike)
	m.Get("/follow_status", oauth2.LoginRequired, handlers.GetFollowStatus)
	m.Get("/update_follow_status", oauth2.LoginRequired, handlers.UpdateFollowStatus)
	m.Get("/feed", oauth2.LoginRequired, handlers.FeedHandler)
	m.Post("/avatar/upload", oauth2.LoginRequired, handlers.UploadAvatar)
	m.Get("/following", oauth2.LoginRequired, handlers.FollowingHandler)
	m.Get("/album/:id", oauth2.LoginRequired, handlers.AlbumHandler)
	m.Post("/load_photo_album", oauth2.LoginRequired, handlers.LoadPhotoAlbum)
	m.Run()
}