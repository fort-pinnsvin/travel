package handlers

import (
	"encoding/json"
	gooauth2 "github.com/golang/oauth2"
	"github.com/martini-contrib/oauth2"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/jmoiron/jsonq"
	"strings"
	"strconv"
	"github.com/fort-pinnsvin/travel/models"
)

func GetCountry(tokens oauth2.Tokens, lat string, lng string) {
	tr := gooauth2.NewTransport(http.DefaultTransport, nil, &gooauth2.Token{AccessToken: tokens.Access()})
	client := http.Client{Transport: tr}
	res, err := client.Get("http://maps.googleapis.com/maps/api/geocode/json?latlng=" + lat + "," + lng)
	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	dat := map[string]interface{}{}
	dec := json.NewDecoder(strings.NewReader(string(robots)))
	dec.Decode(&dat)
	jq := jsonq.NewQuery(dat)
	arrOfObj, _ := jq.ArrayOfObjects("results")
	len := len(arrOfObj)
	country, err := jq.String("results", strconv.Itoa(len - 1), "address_components", "0", "short_name")
	println(country)


	countryStat := &models.Country{}
	models.CountryCollection.FindId(country).One(&countryStat)

	if countryStat.Code == "" {
		countryStat.Code = country
		countryStat.Count = 1
		models.CountryCollection.Insert(&countryStat)
	} else {
		countryStat.Count += 1
		models.CountryCollection.UpdateId(country, countryStat)
	}
}
