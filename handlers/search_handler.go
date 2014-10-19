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
	words := strings.Split(query, " ")
	if len(query) != 0 {
		if strings.Contains(query, " ") {
			if (words[0] == "->"){
				q["title"] = "With me to " + words[1];
			}else{
				q["firstname"] = words[0]
				q["lastname"] = words[1]
			}
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
	posts := []models.Post{}
	if (words[0] == "->"){
		models.PostCollection.Find(q).Limit(1024).Iter().All(&posts)
	}else {
		iter := models.UserCollection.Find(q).Limit(100).Iter()
		if err := iter.All(&users); err == nil {
		}
	}

		if len(query) == 0 {
			query = "[empty query]"
		}
		user := &models.User{}
		user.FirstName = session.Get("first_name").(string)
		user.LastName = session.Get("last_name").(string)
		user.Id = session.Get("auth_id").(string)
		user.Avatar = session.Get("avatar").(string)



		if (words[0] == "->"){
			for i := len(posts) - 1; i >= 0; i-- {
				user := models.User{}
				models.UserCollection.FindId(posts[i].Owner).One(&user)
				posts[i].OwnerUser = user
				posts[i].IsLiked = IsPostLiked(session.Get("auth_id").(string), posts[i].Id)
			}
			rnd.HTML(200, "search_with_me", map[string]interface{}{"user": user, "result": posts, "query": query})
		}else {
			if (len(users) != 0) {
				rnd.HTML(200, "search", map[string]interface{}{"user": user, "result": users, "query": query})
			} else {
				rnd.HTML(200, "search_empty", map[string]interface{}{"user": user, "query": query})
			}
		}
	}

