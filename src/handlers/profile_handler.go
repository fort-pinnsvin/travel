package handlers

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"models"
)

func UserProfile(rnd render.Render, params martini.Params, session sessions.Session) {
	if session.Get("auth_id") != "" {
		id := params["id"]
		user := models.User{}
		user_auth := models.User{}
		err := models.UserCollection.FindId(id).One(&user)
		models.UserCollection.FindId(session.Get("auth_id")).One(&user_auth)
		if err != nil {
			fmt.Println(err)
			rnd.Redirect("/")
			return
		}
		b := ""
		if id == session.Get("auth_id") {
			b = "true"
		}
		rnd.HTML(200, "user", map[string]string{
			"auth_first_name": user_auth.FirstName,
			"auth_last_name": user_auth.LastName,
			"auth_avatar": user_auth.Avatar,
			"auth_id": user_auth.Id,
			"first_name": user.FirstName,
			"last_name": user.LastName,
			"email": user.Email,
			"avatar": user.Avatar,
			"auth_user": b,
			"country": user_auth.Country,
			"birthday": user_auth.Birthday,
			"about": user_auth.About,
		})
	} else {
		rnd.HTML(200, "not_allowed", map[string]string{"error": "Not authorized"})
	}
}
