package handlers

import (
	"encoding/json"
	"github.com/fort-pinnsvin/travel/models"
	gooauth2 "github.com/golang/oauth2"
	"github.com/jmoiron/jsonq"
	"github.com/martini-contrib/oauth2"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"io/ioutil"
	"labix.org/v2/mgo/bson"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
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
	status, _ := jq.String("status")
	if status == "OK" {
		arrOfObj, _ := jq.ArrayOfObjects("results")
		len := len(arrOfObj)
		country, _ := jq.String("results", strconv.Itoa(len-1), "address_components", "0", "short_name")
		name, _ := jq.String("results", strconv.Itoa(len-1), "address_components", "0", "long_name")
		println(name)

		countryStat := &models.Country{}
		models.CountryCollection.FindId(country).One(&countryStat)

		if countryStat.Code == "" {
			countryStat.Code = country
			countryStat.Name = name
			countryStat.Count = 1
			models.CountryCollection.Insert(&countryStat)
		} else {
			countryStat.Count += 1
			models.CountryCollection.UpdateId(country, countryStat)
		}
	}
}

func GetRecommCountry(tokens oauth2.Tokens, res http.ResponseWriter, r *http.Request, session sessions.Session, rnd render.Render) {
	country := []models.Country{}
	query := make(bson.M)
	iter := models.CountryCollection.Find(query).Limit(10).Iter()
	if err := iter.All(&country); err == nil {
		sort.Sort(models.ByCountry(country))
		user := &models.User{}
		user.FirstName = session.Get("first_name").(string)
		user.LastName = session.Get("last_name").(string)
		user.Id = session.Get("auth_id").(string)
		user.Avatar = session.Get("avatar").(string)
		rnd.HTML(200, "advice", map[string]interface{}{
			"user":    user,
			"country": country,
		})
	}
}


func GetLatLngByAddress(tokens oauth2.Tokens, address string) (float64, float64) {
	tr := gooauth2.NewTransport(http.DefaultTransport, nil, &gooauth2.Token{AccessToken: tokens.Access()})
	client := http.Client{Transport: tr}
	res, err := client.Get("http://maps.googleapis.com/maps/api/geocode/json?address=" + strings.Replace(address, " ", "+", -1))
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
	status, _ := jq.String("status")
	if status == "OK" {
		lat, _ := jq.Float("results", "0", "geometry", "location", "lat")
		lng, _ := jq.Float("results", "0", "geometry", "location", "lng")
		return lat, lng
	}
	return 10000, 10000
}
