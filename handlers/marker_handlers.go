package handlers

import (
	"github.com/martini-contrib/sessions"
	"net/http"
	"github.com/martini-contrib/oauth2"
	"labix.org/v2/mgo/bson"
	"github.com/FortPinnsvin/Travel/models"
	"encoding/json"
	"fmt"
)

func GetMarkers(tokens oauth2.Tokens, res http.ResponseWriter, r *http.Request, session sessions.Session) {
	id := ""
	if r.FormValue("id") != "" {
		id = r.FormValue("id")
	} else if session.Get("auth_id") != "" {
		id = session.Get("auth_id").(string)
	}
	markers := []models.Marker{}
	query := make(bson.M)
	query["owner"] = id
	fmt.Printf("%v\n", query)
	fmt.Printf("%v\n", id)
	iter := models.MarkerCollection.Find(query).Limit(1024).Iter()
	if err := iter.All(&markers); err == nil {
		fmt.Printf("%v\n", markers)
		ans, _ := json.Marshal(markers)
		res.Write(ans)
	}
}

func CreateMarker(tokens oauth2.Tokens, res http.ResponseWriter, r *http.Request, session sessions.Session) {
	if session.Get("auth_id") != "" {
		marker := &models.Marker{}
		marker.Id = models.GenerateId()
		marker.Owner = session.Get("auth_id").(string)
		marker.Name = r.FormValue("name")
		marker.Latitude = r.FormValue("lat")
		marker.Longitude = r.FormValue("long")
		models.MarkerCollection.Insert(&marker)

		res.Write([]byte(`{"error": 0, "id": "` + marker.Id + `"}`))
	} else {
		res.Write([]byte(`some errors`))
	}
}
