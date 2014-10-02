package handlers

import (
	"github.com/martini-contrib/sessions"
	"github.com/martini-contrib/render"
	"github.com/go-martini/martini"
	"net/http"
	//"os"
	//"io"
	"fmt"
	"io"
	"os"
)

func AlbumHandler(rnd render.Render, session sessions.Session,  params martini.Params){
	if session.Get("auth_id") != "" {
		id := params["id"]
		fmt.Println(id + "    222222222222");
		rnd.HTML(200, "album_empty", id)
	}
}

func LoadPhotoAlbum(r *http.Request, session sessions.Session){
	if session.Get("auth_id") != "" {

		album_id := r.FormValue("id")

		file, name_file , err := r.FormFile("file") // the FormFile function takes in the POST input id file
		defer file.Close()

		if err != nil {
			fmt.Println(err)
		}

		if os.MkdirAll("album/" +album_id+ "/", 0777) != nil {
			fmt.Println("Unable to create directory for tagfile!")
		}

		out, err := os.Create("album/" +album_id+"/"+ name_file.Filename)

		if err != nil {
			fmt.Println("Unable to create the file for writing. Check your write access privilege")
		}

		defer out.Close()

		_,err = io.Copy(out, file)

		if err != nil {
			fmt.Println(err)
		}
	}
}

