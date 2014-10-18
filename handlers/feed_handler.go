package handlers

import (
	"github.com/fort-pinnsvin/travel/helpfunc"
	"github.com/fort-pinnsvin/travel/models"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"labix.org/v2/mgo/bson"
	"sort"
	"html/template"
)

func FeedHandler(rnd render.Render, params martini.Params, session sessions.Session) {
	if session.Get("auth_id") != "" {
		id := session.Get("auth_id").(string)
		user := helpfunc.GetAuthUser(session)
		allPosts := []models.Post{}
		followingList := []models.FollowEdge{}
		models.PostCollection.Find(bson.M{"owner": id}).All(&allPosts)
		for i := len(allPosts) - 1; i >= 0; i-- {
			allPosts[i].Html = template.HTML(allPosts[i].Text)
			allPosts[i].OwnerUser = user
			allPosts[i].IsLiked = IsPostLiked(session.Get("auth_id").(string), allPosts[i].Id)
		}
		models.FollowCollection.Find(bson.M{"follower": id}).All(&followingList)
		for _, element := range followingList {
			user_id := element.Following
			posts := []models.Post{}
			models.PostCollection.Find(bson.M{"owner": user_id}).All(&posts)
			for i := len(posts) - 1; i >= 0; i-- {
				user := models.User{}
				models.UserCollection.FindId(element.Following).One(&user)
				posts[i].OwnerUser = user
				posts[i].IsLiked = IsPostLiked(session.Get("auth_id").(string), posts[i].Id)
				posts[i].Html = template.HTML(posts[i].Text)
				allPosts = append(allPosts, posts[i])
			}
		}

		sort.Sort(models.ByPost(allPosts))

		rnd.HTML(200, "feed", map[string]interface{}{"user": user, "posts": allPosts, "following": followingList})
	} else {
		rnd.HTML(200, "not_allowed", map[string]string{"error": "Not authorized"})
	}
}
