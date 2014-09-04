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
)

func GetData(tokens oauth2.Tokens)  {
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
	if err := json.Unmarshal(robots, &dat); err == nil {
		user := &models.User{};
		user.Id = string(dat["id"])
		user.FirstName = string(dat["given_name"])
		user.LastName = string(dat["family_name"])
		user.Email = string(dat["link"])
		fmt.Printf("%v\n", user)
	}
}
