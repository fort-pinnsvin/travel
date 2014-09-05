package handlers

import (
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"models"
	"net/http"
)

func Edit(rnd render.Render, session sessions.Session, r *http.Request) {
	if session.Get("auth_id") != "" {
		userData := &models.User{}
		models.UserCollection.FindId(session.Get("auth_id")).One(&userData)
		rnd.HTML(200, "edit_profile", userData)
	}
}
func EditPost(res http.ResponseWriter, session sessions.Session, r *http.Request){
	if session.Get("auth_id") != "" {
		user := &models.User{}
		models.UserCollection.FindId(session.Get("auth_id")).One(&user)
		firstName := r.FormValue("first_name");
		lastName := r.FormValue("last_name");
		email := r.FormValue("email");
		user.FirstName = firstName
		user.LastName = lastName
		user.Email = email
		session.Set("first_name", user.FirstName)
		session.Set("last_name", user.LastName)
		models.UserCollection.UpdateId(session.Get("auth_id"),user)
		res.Write([]byte(`{"error":0, "url":"/"}`))
	}
}
