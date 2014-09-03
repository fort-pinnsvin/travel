package handlers

import (
	"github.com/martini-contrib/render"
	"net/http"
	"github.com/martini-contrib/sessions"
	"models"
	"fmt"
)

func Root(rnd render.Render, r *http.Request, session sessions.Session) {
	values := map[string]interface{}{}
	if session.Get("auth") != "" && session.Get("auth_id") != "" {
		user := models.User{}
		err := models.UserCollection.FindId(session.Get("auth_id")).One(&user)
		if err != nil {
			fmt.Println(err)
			rnd.Redirect("/")
			return
		}
		values["id"] = user.FirstName + " " + user.LastName
	} else {
		values["id"] = "guest"
	}
	rnd.HTML(200, "home", values)
}
