package handlers

import (
	"fmt"
	"github.com/fort-pinnsvin/travel/models"
	"github.com/martini-contrib/oauth2"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"net/http"
)

func Edit(rnd render.Render, session sessions.Session, r *http.Request) {
	if session.Get("auth_id") != "" {
		userData := &models.User{}
		models.UserCollection.FindId(session.Get("auth_id")).One(&userData)
		rnd.HTML(200, "edit_profile", userData)
	}
}

func EditPost(tokens oauth2.Tokens, res http.ResponseWriter, session sessions.Session, r *http.Request) {
	if session.Get("auth_id") != "" {
		user := &models.User{}
		models.UserCollection.FindId(session.Get("auth_id")).One(&user)
		firstName := r.FormValue("first_name")
		lastName := r.FormValue("last_name")
		email := r.FormValue("email")
		country := r.FormValue("country")
		birthday := r.FormValue("birthday")
		about := r.FormValue("about")
		lang := r.FormValue("lang")
		lat, lng := GetLatLngByAddress(tokens, country)

		edit_user := models.User{user.Id, firstName, lastName, email, user.Avatar, birthday, country, user.Status, about, lang, lat, lng}

		session.Set("first_name", user.FirstName)
		session.Set("last_name", user.LastName)
		session.Set("avatar", user.Avatar)
		session.Set("lang", lang)
		models.UserCollection.UpdateId(session.Get("auth_id"), edit_user)
		res.Write([]byte(fmt.Sprintf(`{"url": "%s", "error": 0}`, "/edit?ok")))
	} else {
		res.Write([]byte(fmt.Sprintf(`{"url": "%s", "error": 1}`, "/edit?error")))
	}
}
