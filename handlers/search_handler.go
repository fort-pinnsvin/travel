package handlers

import (
	"github.com/fort-pinnsvin/travel/models"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"labix.org/v2/mgo/bson"
	"net/http"
	"strings"
)

func Search(rnd render.Render, params martini.Params, session sessions.Session, r *http.Request) {
	query := r.FormValue("q")
	users := []models.User{}
	q := make(bson.M)
	if len(query) != 0 {
		if strings.Contains(query, " ") {
			words := strings.Split(query, " ")
			q["firstname"] = words[0]
			q["lastname"] = words[1]
		} else {
			countLastName, _ := models.UserCollection.Find(bson.M{"lastname": query}).Count()
			countFirstName, _ := models.UserCollection.Find(bson.M{"firstname": query}).Count()
			if countFirstName > countLastName {
				q["firstname"] = query
			} else {
				q["lastname"] = query
			}
		}
	}
	iter := models.UserCollection.Find(q).Limit(100).Iter()
	if err := iter.All(&users); err == nil {
		user := &models.User{}
		user.FirstName = session.Get("first_name").(string)
		user.LastName = session.Get("last_name").(string)
		user.Id = session.Get("auth_id").(string)
		user.Avatar = session.Get("avatar").(string)
		if len(query) == 0 {
			query = "[empty query]"
		}
		if len(users) != 0 {
			rnd.HTML(200, "search", map[string]interface{}{"user": user, "result": users, "query": query})
		} else {
			rnd.HTML(200, "search_empty", map[string]interface{}{"user": user, "query": query})
		}
	}
}
