package handlers

import (
	"fmt"
	"github.com/fort-pinnsvin/travel/models"
	"github.com/fort-pinnsvin/travel/utils"
	"github.com/martini-contrib/oauth2"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"net/http"
	"time"
)

func Root(tokens oauth2.Tokens, rnd render.Render, r *http.Request, session sessions.Session) {
	if tokens.IsExpired() {
		rnd.HTML(200, "home", nil)
	} else {
		user := &models.User{}
		user.FirstName = session.Get("first_name").(string)
		user.LastName = session.Get("last_name").(string)
		user.Id = session.Get("auth_id").(string)
		user.Avatar = session.Get("avatar").(string)
		if r.FormValue("address") != "" {
			println(r.FormValue("address"))
			lat, lng := GetLatLngByAddress(tokens, r.FormValue("address"))
			marker := &models.Marker{}
			marker.Id = models.GenerateId()
			marker.Owner = session.Get("auth_id").(string)
			marker.Name = r.FormValue("address")
			marker.Latitude = fmt.Sprintf("%0.8f", lat)
			marker.Longitude = fmt.Sprintf("%0.8f", lng)
			marker.Description = ""
			marker.FullAddress = "http://placehold.it/250x130"
			marker.Date = time.Now().Format(models.Layout)
			marker.Nano = time.Now().Unix()
			marker.Country = GetCountry(tokens, marker.Latitude, marker.Longitude)
			println("marker.Country", marker.Country)
			models.MarkerCollection.Insert(&marker)
			// Add marker to feed
			new_post := models.Post{}
			new_post.Id = models.GenerateId()
			new_post.Owner = session.Get("auth_id").(string)
			new_post.Text = `<img src="http://placehold.it/250x130"/><br/>Watch it <a href="` +
				"//" + utils.GetValue("WWW", "localhost:3000") + "/album/" + marker.Id + "/" + `">here</a>.`
			new_post.Title = "I create New Album!"
			new_post.Date = time.Now().Format(models.Layout)
			new_post.Nano = time.Now().Unix()
			models.PostCollection.Insert(&new_post)

			rnd.Redirect("/")
		}
		rnd.HTML(200, "home_user", user)
	}
}
