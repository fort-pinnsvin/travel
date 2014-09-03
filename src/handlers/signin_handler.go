package handlers

import (
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"models"
	"net/http"
)

func PostSignIn(rnd render.Render, r *http.Request, session sessions.Session) {
	users := []models.User{}
	models.UserCollection.Find(nil).All(&users)
	for _, user := range users {
		if user.Email == r.FormValue("email_signin") && user.Password == r.FormValue("password_signin") {
			session.Set("auth", "OK")
			session.Set("auth_id", user.Id)
			rnd.Redirect("/user/" + user.Id)
		}
	}
}

func SignInForm(rnd render.Render) {
	rnd.HTML(200, "signin", nil)
}
