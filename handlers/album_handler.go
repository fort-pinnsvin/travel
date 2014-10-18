package handlers

import (
	"fmt"
	"github.com/fort-pinnsvin/travel/helpfunc"
	"github.com/fort-pinnsvin/travel/models"
	"github.com/fort-pinnsvin/travel/utils"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"labix.org/v2/mgo/bson"
	"net/http"
	"os"
	"path/filepath"
	"time"
	"strings"
)

func AlbumHandler(rnd render.Render, session sessions.Session, params martini.Params) {
	if session.Get("auth_id") != "" {
		// Id album
		id := params["id"]
		// Get markerAlbum by Id
		markerAlbum := models.Marker{}
		models.MarkerCollection.FindId(id).One(&markerAlbum)

		user := helpfunc.GetAuthUser(session)

		// Check, who owner this album
		isOwner := false
		if markerAlbum.Owner == user.Id {
			isOwner = true
		}

		// Get info about Album
		infoAlbum := models.Marker{}
		models.MarkerCollection.FindId(id).One(&infoAlbum)

		// Get all photos from this album
		allPhoto := []models.Photo{}
		query := make(bson.M)
		query["albumid"] = id
		iter := models.PhotoCollection.Find(query).Limit(1024).Iter()
		_ = iter.All(&allPhoto)

		rnd.HTML(200, "album_empty", map[string]interface{}{
			"id":         id,
			"user":       user,
			"owner":      isOwner,
			"photos":     allPhoto,
			"info_album": infoAlbum,
			"auth_user" : true,
		})
	}
}

func LoadPhotoAlbum(r *http.Request, session sessions.Session, rnd render.Render) string {
	if session.Get("auth_id") != "" {

		album_id := r.FormValue("id")
		fmt.Println(album_id + "------9999999")

		file, name_file, err := r.FormFile("file") // the FormFile function takes in the POST input id file
		defer file.Close()

		if err != nil {
			fmt.Println(err)
		}

		if os.MkdirAll("assets/album/"+album_id+"/", 0777) != nil {
			fmt.Println("Unable to create directory for tagfile!")
		}
		file_id := models.GenerateId()
		extension := filepath.Ext(name_file.Filename)
		out, err := os.Create("assets/album/" + album_id + "/" + file_id + extension)

		photo := models.Photo{}

		if err != nil {
			fmt.Println("Unable to create the file for writing. Check your write access privilege")
		}

		defer out.Close()

		_, err = io.Copy(out, file)

		if err != nil {
			fmt.Println(err)
		}
		if isValidImage("assets/album/" + album_id + "/" + file_id + extension) {
			photo.AlbumId = album_id
			photo.Name = file_id + extension
			models.PhotoCollection.Insert(&photo)

			// Add marker to feed
			new_post := models.Post{}
			new_post.Id = models.GenerateId()
			new_post.Owner = session.Get("auth_id").(string)
			new_post.Text = `I upload new photo to <a href="` +
				"//" + utils.GetValue("WWW", "localhost:3000") + "/album/" + album_id + "/" +
				`">album</a><br/><img src="` +
				"//" + utils.GetValue("WWW", "localhost:3000") + "/album/" + album_id + "/" + file_id + extension + `" height="130px">`
			new_post.Title = "I create New Album!"
			new_post.Date = time.Now().Format(models.Layout)
			new_post.Nano = time.Now().Unix()
			models.PostCollection.Insert(&new_post)

			marker := &models.Marker{}
			models.MarkerCollection.FindId(album_id).One(&marker)
			// Adress last photo
			marker.FullAddress = "//" + utils.GetValue("WWW", "localhost:3000") + "/album/" + album_id + "/" + file_id + extension
			models.MarkerCollection.UpdateId(album_id, marker)
			rnd.Redirect("/album/" + album_id)
			return "ok"
		} else {
			os.Remove("assets/album/" + album_id + "/" + file_id + extension)
			rnd.Redirect("/album/" + album_id)
			return "error"
		}
	}
	return ""
}

func isValidImage(filename string) bool {
	file, err := os.Open(filename)
	if err != nil {
		return false
	}
	_, _, err_img := image.DecodeConfig(file)
	return (err_img == nil)
}

func AlbumSettingsHandler(rnd render.Render, session sessions.Session, r *http.Request, params martini.Params) {
	if session.Get("auth_id") != "" {
		id := params["id"]
		userData := &models.User{}
		marker := &models.Marker{}
		models.UserCollection.FindId(session.Get("auth_id")).One(&userData)
		models.MarkerCollection.FindId(id).One(&marker)
		rnd.HTML(200, "album_settings", map[string]interface{}{
			"Id":          userData.Id,
			"Avatar":      userData.Avatar,
			"FirstName":   userData.FirstName,
			"LastName":    userData.LastName,
			"title":       marker.Name,
			"description": marker.Description,
			"album_id":    marker.Id,
		})
	}
}

func AlbumSettingsSaveHandler(res http.ResponseWriter, session sessions.Session, r *http.Request, params martini.Params) {
	if session.Get("auth_id") != "" {
		id := params["id"]
		marker := &models.Marker{}
		models.MarkerCollection.FindId(id).One(&marker)
		marker.Name = r.FormValue("title")
		marker.Description = r.FormValue("description")
		models.MarkerCollection.UpdateId(id, marker)
		res.Write([]byte(`{"error": 0}`))
	} else {
		res.Write([]byte(`{"error": 1}`))
	}
}

func RemovePhoto(r *http.Request, session sessions.Session) {
	if session.Get("auth_id") != "" {
		photo_name := r.FormValue("name_photo")
		query := make(bson.M)
		query["name"] = photo_name
		photo := models.Photo{}
		models.PhotoCollection.Find(query).One(&photo)
		models.PhotoCollection.Remove(query)
		os.Remove("assets/album/" + photo.AlbumId + "/" + photo.Name)
		// Remove photo from feeds
		posts := []models.Post{}
		query = make(bson.M)
		query["owner"] = session.Get("auth_id").(string)
		iter := models.PostCollection.Find(query).Limit(1024).Iter()
		if err := iter.All(&posts); err == nil {
			for  i := 0; i < len(posts); i ++ {
				if strings.Contains(posts[i].Text, photo.Name) {
					models.PostCollection.RemoveId(posts[i].Id)
				}
			}
		}
	}
}

func AlbumDeleteHandler(res http.ResponseWriter, session sessions.Session, r *http.Request, params martini.Params, rnd render.Render) {
	if session.Get("auth_id") != "" {
		id := params["id"]
		// Remove files
		os.RemoveAll("assets/album/" + id + "/")
		query := make(bson.M)
		query["albumid"] = id
		models.PhotoCollection.RemoveAll(query)
		marker := &models.Marker{}
		models.MarkerCollection.FindId(id).One(&marker)
		models.MarkerCollection.RemoveId(id)
		// TODO
		DecrementCountryStat(marker.Country)
		// Delete posts from feed
		posts := []models.Post{}
		query = make(bson.M)
		query["owner"] = session.Get("auth_id").(string)
		iter := models.PostCollection.Find(query).Limit(1024).Iter()
		if err := iter.All(&posts); err == nil {
			for  i := 0; i < len(posts); i ++ {
				if strings.Contains(posts[i].Text, id) {
					models.PostCollection.RemoveId(posts[i].Id)
				}
			}
		}
		println("OK")
		rnd.Redirect("/")
	}
}
