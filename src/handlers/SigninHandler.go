package handlers

import (
	"github.com/martini-contrib/render"
	"net/http"
	"models"
	"github.com/martini-contrib/sessions"
)


func SigninHandler(rnd render.Render, r *http.Request,session sessions.Session){
	users := []models.UserDocument{}
	models.UserCollection.Find(nil).All(&users)
	for _,user := range users {
		if user.Email == r.FormValue("email_signin") && user.Password == r.FormValue("password_signin") {
			session.Set("auth","OK");
			session.Set("auth_id", user.Id)
			rnd.Redirect("/user/" + user.Id)
		}
	}
}


