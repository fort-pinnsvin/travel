package handlers

import (
	"github.com/martini-contrib/render"
	"models"
	"github.com/go-martini/martini"
	"fmt"
)


func UserHandler(rnd render.Render, params martini.Params){
	id := params["id"]
	user := models.UserDocument{}
	err := models.UserCollection.FindId(id).One(&user)

	if err != nil{
		fmt.Println(err)
		rnd.Redirect("/")
		return
	}

	userr := models.UserDocument{user.Id, user.FirstName, user.LastName, user.Email, user.Password}

	rnd.HTML(200,"user",userr)

}
