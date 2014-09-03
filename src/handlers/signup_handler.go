package handlers

import (
	"fmt"
	"github.com/martini-contrib/render"
	"models"
	"net/http"
)

func SignUpForm(rnd render.Render) {
	rnd.HTML(200, "signup", nil)
}

func PostSignUp(rnd render.Render, r *http.Request) {
	firstName := r.FormValue("first_name")
	lastName := r.FormValue("last_name")
	email := r.FormValue("email")
	confirmEmail := r.FormValue("confirm_email")
	password := r.FormValue("password")
	confirmPassword := r.FormValue("confirm_password")
	if (email != confirmEmail) || (password != confirmPassword) {
		fmt.Print("emails and (or) passwords are not equal!")
		rnd.Redirect("/")
	}
	newUser := &models.User{models.GenerateId(), firstName, lastName, email, password}
	models.UserCollection.Insert(newUser)
	rnd.Redirect("/")
}
