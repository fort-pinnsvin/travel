package handlers

import (
	"net/http"
	"log"
	"io/ioutil"
	"fmt"
	"github.com/martini-contrib/oauth2"
	gooauth2 "github.com/golang/oauth2"
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
	fmt.Printf("%s", robots)
}
