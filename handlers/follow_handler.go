package handlers

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"github.com/fort-pinnsvin/travel/models"
	"labix.org/v2/mgo/bson"
	"fmt"
)

func FollowingHandler(rnd render.Render, params martini.Params, session sessions.Session) {
	if session.Get("auth_id") != "" {
		my_id := session.Get("auth_id").(string)
		user := &models.User{}
		user.FirstName = session.Get("first_name").(string)
		user.LastName = session.Get("last_name").(string)
		user.Id = session.Get("auth_id").(string)
		user.Avatar = session.Get("avatar").(string)
		follower := []models.FollowEdge{}
		models.FollowCollection.Find(bson.M{"follower": my_id}).All(&follower)
		fmt.Println(follower)
		follower_user := []models.User{}
		//allFollower := []models.User{}

		for _, element := range follower {
			user_id := element.Following
			all_user := []models.User{}
			models.UserCollection.Find(bson.M{"_id": user_id}).All(&all_user)
			for i := len(all_user) - 1; i >= 0; i-- {
				user := models.User{}
				models.UserCollection.FindId(element.Following).One(&user)
				follower_user = append(follower_user, user)
			}
		}

		following := []models.FollowEdge{}
		models.FollowCollection.Find(bson.M{"following": my_id}).All(&following)
		following_user := []models.User{}
		fmt.Println(following)
		for _, element := range following {
			user_id := element.Follower
			all_user := []models.User{}
			models.UserCollection.Find(bson.M{"_id": user_id}).All(&all_user)

			for i := len(all_user) - 1; i >= 0; i-- {
				user := models.User{}
				models.UserCollection.FindId(element.Follower).One(&user)
				following_user = append(following_user, user)
			}
		}
		rnd.HTML(200, "follower", map[string]interface{}{"user": user, "follower":follower_user, "following":following_user})
	}
}
