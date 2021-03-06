package main

import (
	"github.com/fort-pinnsvin/travel/handlers"
	"github.com/fort-pinnsvin/travel/models"
	"github.com/fort-pinnsvin/travel/utils"
	"github.com/go-martini/martini"
	gooauth2 "github.com/golang/oauth2"
	"github.com/martini-contrib/oauth2"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
)

func main() {
	models.ConnectToDataBase()
	store := sessions.NewCookieStore([]byte("yLiCQYG7CAflDavqGH461IO0MHp7TEbpg6TwHBWdJzNwYod1i5ZTbrIF5bEoO3oP"))
	m := martini.Classic()
	m.Use(martini.Static("assets"))
	m.Use(sessions.Sessions("auth", store))
	m.Use(oauth2.Google(&gooauth2.Options{
		ClientID:     utils.GetValue("CLIENT_ID", "903364406910-m1b4j2vjkfd3qj1npusv6p2qk38fqb3q"),
		ClientSecret: utils.GetValue("CLIENT_SECRET", "iofaFDfJuJRkPTjPu4NuHx61"),
		RedirectURL:  utils.GetValue("REDIRECT", "http://localhost:3000/oauth2callback"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile"},
	}))
	m.Use(render.Renderer(render.Options{
		Directory:  "templates",
		Extensions: []string{".tmpl", ".html"},
		Charset:    "UTF-8",
		IndentJSON: true,
	}))

	m.Get("/", handlers.Root)
	m.Get("/weather", oauth2.LoginRequired, handlers.Weather)
	m.Get("/user/:id", oauth2.LoginRequired, handlers.UserProfile)
	m.Get("/signin", oauth2.LoginRequired, handlers.GetData)
	m.Get("/edit", oauth2.LoginRequired, handlers.Edit)
	m.Get("/search", oauth2.LoginRequired, handlers.Search)
	m.Get("/markers", oauth2.LoginRequired, handlers.GetMarkers)
	m.Get("/markers/create", oauth2.LoginRequired, handlers.CreateMarker)
	m.Get("/markers/update", oauth2.LoginRequired, handlers.UpdateMarkerLocation)
	m.Get("/css/bootstrap.min.css", handlers.LoadTheme)
	m.Get("/follow_status", oauth2.LoginRequired, handlers.GetFollowStatus)
	m.Get("/update_follow_status", oauth2.LoginRequired, handlers.UpdateFollowStatus)
	m.Get("/feed", oauth2.LoginRequired, handlers.FeedHandler)
	m.Get("/friends", oauth2.LoginRequired, handlers.FollowingHandler)
	m.Get("/album_list", oauth2.LoginRequired, handlers.AlbumListHandler)
	m.Get("/advice/country", oauth2.LoginRequired, handlers.GetRecommCountry)
	m.Get("/album/:id", oauth2.LoginRequired, handlers.AlbumHandler)
	m.Get("/album/:id/settings", oauth2.LoginRequired, handlers.AlbumSettingsHandler)
	m.Get("/album/:id/delete", oauth2.LoginRequired, handlers.AlbumDeleteHandler)
	m.Get("/user/:id/home", oauth2.LoginRequired, handlers.GetHomePosition)
	m.Get("/mini_blog/:id",oauth2.LoginRequired,handlers.MiniBlogHandler)
	m.Get("/mini_blog_list/:id",oauth2.LoginRequired,handlers.MiniBlogListHandler)
	m.Get("/mini_blog/:id/edit",oauth2.LoginRequired,handlers.MiniBlogEdit)
	m.Get("/routes/editor", oauth2.LoginRequired, handlers.RouteEditor)
	m.Get("/routes/create", oauth2.LoginRequired, handlers.CreateRoute)
	m.Get("/routes", oauth2.LoginRequired, handlers.RouteHandler)
	m.Get("/routes/:id", oauth2.LoginRequired, handlers.RouteViewer)
	m.Get("/routes/:id/delete", oauth2.LoginRequired, handlers.RemoveRoute)
	m.Post("/album/:id/settings", oauth2.LoginRequired, handlers.AlbumSettingsSaveHandler)
	m.Post("/load_photo_album", oauth2.LoginRequired, handlers.LoadPhotoAlbum)
	m.Post("/update", oauth2.LoginRequired, handlers.EditPost)
	m.Post("/set_theme", handlers.SetTheme)
	m.Post("/save_post", oauth2.LoginRequired, handlers.SavePost)
	m.Post("/remove_post", oauth2.LoginRequired, handlers.RemovePost)
	m.Post("/add_like", oauth2.LoginRequired, handlers.AddLike)
	m.Post("/avatar/upload", oauth2.LoginRequired, handlers.UploadAvatar)
	m.Post("/remove_photo",oauth2.LoginRequired, handlers.RemovePhoto)
	m.Post("/blog_create",oauth2.LoginRequired, handlers.CreateMiniBlog)
	m.Post("/save_post_miniblog",oauth2.LoginRequired, handlers.SavePostMiniblog)
	m.Post("/remove_post_miniblog",oauth2.LoginRequired, handlers.RemovePostMiniblog)
	m.Post("/mini_blog/:id/edit", oauth2.LoginRequired, handlers.SaveEditBlog)

	m.Run()
}
