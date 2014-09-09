package handlers

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"labix.org/v2/mgo/bson"
	"models"
	"sort"
)

func FeedHandler(rnd render.Render, params martini.Params, session sessions.Session) {
	if session.Get("auth_id") != "" {
		id := session.Get("auth_id").(string)
		user := models.User{}
		user.FirstName = session.Get("first_name").(string)
		user.LastName = session.Get("last_name").(string)
		user.Id = session.Get("auth_id").(string)
		user.Avatar = session.Get("avatar").(string)
		allPosts := []models.Post{}
		followingList := []models.FollowEdge{}
		models.PostCollection.Find(bson.M{"owner": id}).All(&allPosts)
		for i := len(allPosts) - 1; i >= 0; i-- {
			allPosts[i].OwnerUser = user
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
				allPosts = append(allPosts, posts[i])
			}
		}
		sort.Sort(models.ByPost(allPosts))

		rnd.HTML(200, "feed", map[string]interface{}{"user": user, "posts": allPosts, "following": followingList})
	} else {
		rnd.HTML(200, "not_allowed", map[string]string{"error": "Not authorized"})
	}
}
