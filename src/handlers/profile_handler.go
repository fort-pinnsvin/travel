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
		err := models.UserCollection.FindId(id).One(&user)

		if err != nil {
			fmt.Println(err)
			rnd.Redirect("/")
			return
		}

		rnd.HTML(200, "user", user)
	} else {
		rnd.HTML(200, "not_allowed", map[string]string{"error": "Not authorized"})
	}
}
