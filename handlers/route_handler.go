package handlers

import (
	"github.com/fort-pinnsvin/travel/models"
	"github.com/martini-contrib/oauth2"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"net/http"
	"github.com/jmoiron/jsonq"
	"encoding/json"
	"strings"
	"strconv"
	"fmt"
	"labix.org/v2/mgo/bson"
	"github.com/go-martini/martini"
)

func RouteEditor(tokens oauth2.Tokens, rnd render.Render, r *http.Request, session sessions.Session) {
	if tokens.IsExpired() {
		rnd.HTML(200, "home", nil)
	} else {
		user := &models.User{}
		user.FirstName = session.Get("first_name").(string)
		user.LastName = session.Get("last_name").(string)
		user.Id = session.Get("auth_id").(string)
		user.Avatar = session.Get("avatar").(string)
		rnd.HTML(200, "route_editor", user)
	}
}

func CreateRoute(tokens oauth2.Tokens, res http.ResponseWriter, r *http.Request, session sessions.Session) {
	if session.Get("auth_id") != "" {
		name := r.FormValue("name")
		desc := r.FormValue("description")
		r := r.FormValue("route")
		lats := []float64{};
		long := []float64{};
		dat := map[string]interface{}{}
		dec := json.NewDecoder(strings.NewReader(r))
		dec.Decode(&dat)
		jq := jsonq.NewQuery(dat)
		len, _ := jq.Int("length")
		for i := 0; i < len; i ++ {
			k, _ := jq.Float("j", strconv.Itoa(i), "k")
			b, _ := jq.Float("j", strconv.Itoa(i), "B")
			lats = append(lats, k)
			long = append(long, b)
		}
		route := &models.Route{}
		route.Id = models.GenerateId()
		route.Name = name
		route.Desc = desc
		route.Lat = lats
		route.Long = long
		models.RouteCollection.Insert(&route)
		res.Write([]byte(fmt.Sprintf(`{"error":0}`)))
		return
	}
	res.Write([]byte(fmt.Sprintf(`{"error":1}`)))
}

func RouteHandler(tokens oauth2.Tokens, rnd render.Render, r *http.Request, session sessions.Session) {
	if session.Get("auth_id") != "" {
		user := &models.User{}

		user.FirstName = session.Get("first_name").(string)
		user.LastName = session.Get("last_name").(string)
		user.Id = session.Get("auth_id").(string)
		user.Avatar = session.Get("avatar").(string)
		routes := []models.Route{}
		models.RouteCollection.Find(bson.M{}).All(&routes)
		rnd.HTML(200, "routes", map[string]interface{}{"user": user, "routes": routes,})
	}
}

func RouteViewer(tokens oauth2.Tokens, rnd render.Render, r *http.Request, session sessions.Session, params martini.Params) {
	if session.Get("auth_id") != "" {
		user := &models.User{}

		user.FirstName = session.Get("first_name").(string)
		user.LastName = session.Get("last_name").(string)
		user.Id = session.Get("auth_id").(string)
		user.Avatar = session.Get("avatar").(string)
		route := &models.Route{}
		models.RouteCollection.FindId(params["id"]).One(&route)
		rnd.HTML(200, "route_viewer", map[string]interface{}{"user": user, "Route": route,})
	}
}
