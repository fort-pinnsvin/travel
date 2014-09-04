package test

import (
	"fmt"
	"github.com/martini-contrib/render"
	"models"
	"net/http"
)

func SignUpForm(rnd render.Render) {
	rnd.HTML(200, "signup", nil)
}

func PostSignUp(r *http.Request, res http.ResponseWriter) {
	firstName := r.FormValue("first_name")
	lastName := r.FormValue("last_name")
	email := r.FormValue("email")
	confirmEmail := r.FormValue("confirm_email")
	password := r.FormValue("password")
	confirmPassword := r.FormValue("confirm_password")
	models.UserCollection.Find
	if (email != confirmEmail) || (password != confirmPassword) {
		fmt.Print("emails and (or) passwords are not equal!")
		res.Write([]byte(`{"error":1, "url":"/"}`))
	}
	newUser := &models.User{models.GenerateId(), firstName, lastName, email, password}
	models.UserCollection.Insert(newUser)

	res.Write([]byte(`{"error":0, "url":"/"}`))
}
