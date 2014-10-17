package handlers

import (
	"fmt"
	"github.com/martini-contrib/sessions"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"net/http"
	"os"
	"github.com/fort-pinnsvin/travel/models"
	"path/filepath"
	"github.com/fort-pinnsvin/travel/utils"
	"github.com/martini-contrib/render"
)

func UploadAvatar(session sessions.Session, w http.ResponseWriter, r *http.Request, rnd render.Render) {
	if session.Get("auth_id") != "" {
		id := session.Get("auth_id").(string)
		file, header, err := r.FormFile("file") // the FormFile function takes in the POST input id file
		defer file.Close()

		if err != nil {
			rnd.Redirect("/edit?fail")
			return
		}

		if os.MkdirAll("assets/avatar/"+id+"/", 0777) != nil {
			panic("Unable to create directory for tagfile!" + err.Error())
		}
		file_id := models.GenerateId()
		extension := filepath.Ext(header.Filename)
		out, err := os.Create("assets/avatar/" + id + "/" + file_id + extension)
		if err != nil {
			rnd.Redirect("/edit?fail")
			return
		}

		defer out.Close()

		// write the content from POST to the file
		_, err = io.Copy(out, file)
		if err != nil {
			fmt.Fprintln(w, err)
		}

		filename := "assets/avatar/" + id + "/" + file_id + extension
		width, height := getImageDimension(filename)
		if width != height {
			os.Remove(filename)
			rnd.Redirect("/edit?avatar=not_square")
			return
		}

		user := &models.User{}
		models.UserCollection.FindId(session.Get("auth_id")).One(&user)
		user.Avatar = "//" +utils.GetValue("WWW", "localhost:3000") + "/avatar/" + id + "/" + file_id + extension
		models.UserCollection.UpdateId(session.Get("auth_id"), user)
		session.Set("avatar",  "//" +utils.GetValue("WWW", "localhost:3000") + "/avatar/" + id + "/" + file_id + extension)
		rnd.Redirect("/edit?avatar=ok")
		return
	}
	rnd.Redirect("/edit?avatar=fail")
}

func getImageDimension(imagePath string) (int, int) {
	file, err := os.Open(imagePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return -1, -2
	}

	img, _, err := image.DecodeConfig(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", imagePath, err)
		return -1, -2
	}
	return img.Width, img.Height
}
