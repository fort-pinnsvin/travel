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
	"fmt"
)

func GetCountry(tokens oauth2.Tokens, lat string, lng string) string {
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
		return country
	}
	return ""
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
			"travelers": GetBestUsers(),
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

func DecrementCountryStat(code string) {
	fmt.Printf("DecrementCountryStat(code = [%#v])\n", code)
	if (code != "") {
		countryStat := &models.Country{}
		models.CountryCollection.FindId(code).One(&countryStat)
		if countryStat.Code != "" {
			if countryStat.Count <= 1 {
				models.CountryCollection.RemoveId(code)
			} else {
				countryStat.Count -= 1
				models.CountryCollection.UpdateId(code, countryStat)
			}
		}
	}
}

func GetAddress(tokens oauth2.Tokens, lat string, lng string, def string) string {
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
		address, err := jq.String("results", "6", "formatted_address")
		println("address:", address)
		if address == "" || err != nil {
			return def
		}
		return address
	}
	return def
}

func GetBestUsers() []models.User {
	q := make(bson.M)
	users := []models.User{}
	iter := models.UserCollection.Find(q).Limit(1024).Iter()
	if err := iter.All(&users); err == nil {
		for i := 0; i < len(users); i ++ {
			users[i].Points = GetRate(users[i].Id)
		}
	}
	sort.Sort(models.ByUser(users))
	if len(users) < 10 {
		return users
	} else {
		return users[0:10]
	}
}
