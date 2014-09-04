package handlers

import (
	"net/http"
	"log"
	"io/ioutil"
	"fmt"
	"github.com/martini-contrib/oauth2"
	gooauth2 "github.com/golang/oauth2"
	"encoding/json"
	"models"
	"github.com/martini-contrib/sessions"
	"github.com/martini-contrib/render"
)

func GetData(session sessions.Session, tokens oauth2.Tokens, rnd render.Render) {
	tr := gooauth2.NewTransport(http.DefaultTransport, nil, &gooauth2.Token{AccessToken: tokens.Access()})
	client := http.Client{Transport: tr}
	res, err := client.Get("https://www.googleapis.com/oauth2/v1/userinfo?alt=json")
	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	var dat map[string]string
	user := &models.User{};
	if err := json.Unmarshal(robots, &dat); err == nil {
		user.Id = string(dat["id"])
		user.FirstName = string(dat["given_name"])
		user.LastName = string(dat["family_name"])
		user.Email = string(dat["link"])
	}
	findUser := &models.User{};
	models.UserCollection.FindId(user.Id).One(&findUser)
	fmt.Println(findUser)

	if findUser.Id == "" {
		models.UserCollection.Insert(&user)
		session.Set("auth_id", user.Id)
	}else {
		models.UserCollection.UpdateId(user.Id, &user)
		session.Set("auth_id", user.Id)
	}
	rnd.Redirect("/user/" + user.Id)
}
