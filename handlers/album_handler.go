package handlers

import (
	"github.com/martini-contrib/sessions"
	"github.com/martini-contrib/render"
	"github.com/go-martini/martini"
	"github.com/fort-pinnsvin/travel/helpfunc"
	"net/http"
	"fmt"
	"io"
	"os"
	"github.com/fort-pinnsvin/travel/models"
)

func AlbumHandler(rnd render.Render, session sessions.Session,  params martini.Params){
	if session.Get("auth_id") != "" {
		// Id album
		id := params["id"]
		// Get markerAlbum by Id
		markerAlbum := models.Marker{};
		models.MarkerCollection.FindId(id).One(&markerAlbum)

		user := helpfunc.GetAuthUser(session)

		// Check, who owner this album
		checkOwner := false
		if (markerAlbum.Owner == user.Id){
			checkOwner = true
		}

		rnd.HTML(200, "album_empty", map[string]interface {}{
			"id" : id,
			"user" : user,
			"owner" : checkOwner,
		})
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

