package handlers

import (
	"github.com/martini-contrib/render"
	"net/http"
	"travel/models"
)

func SignupHandler(rnd render.Render, r *http.Request){
	firstName := r.FormValue("first_name")
	lastName := r.FormValue("last_name")
	email_first := r.FormValue("email")
	email_second := r.FormValue("confirm_email")
	pass_first := r.FormValue("password")
	pass_second := r.FormValue("confirm_password")
	if (email_first != email_second) {
		rnd.Redirect("/")
	}
	if pass_first != pass_second{
		rnd.Redirect("/")
	}
	newUser := &models.UserDocument{models.GenerateId(),firstName,lastName,email_first,pass_first}
	models.UserCollection.Insert(newUser)
	rnd.Redirect("/")
}

func SigninHandler(rnd render.Render, r *http.Request){
	users := []models.UserDocument{}
	UserCollection.Find(nil).All(&users)
	for _,user := range users {
		if user.Email == r.FormValue("email_signin") && user.Password == r.FormValue("password_signin") {
			rnd.Redirect("/" + user.FirstName)
		}
	}
}


