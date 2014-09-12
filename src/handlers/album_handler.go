package handlers

import (
	"github.com/martini-contrib/sessions"
	"github.com/martini-contrib/render"
	"github.com/go-martini/martini"
	"net/http"
	"os"
	"io"
	"fmt"
)

func AlbumHandler(rnd render.Render, session sessions.Session,  params martini.Params){
	if session.Get("auth_id") != "" {
		id := params["id"]
		rnd.HTML(200, "album_empty", id)
	}
}

func LoadPhotoAlbum(r *http.Request, session sessions.Session){
	if session.Get("auth_id") != "" {
		file, name_file , _ := r.FormFile("file") // the FormFile function takes in the POST input id file
		defer file.Close()
		out, _ := os.Create("album/" + session.Get("auth_id").(string) + "/" + name_file)

		defer out.Close()
		fmt.Println(out)
		_, _ = io.Copy(out, file)
		fmt.Println("OK 3 !!!")
	}
}

