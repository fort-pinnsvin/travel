package handlers

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"labix.org/v2/mgo/bson"
	"models"
	"net/http"
	"strconv"
	"time"
)

func UserProfile(rnd render.Render, params martini.Params, session sessions.Session) {
	if session.Get("auth_id") != "" {
		id := params["id"]
		user := models.User{}
		user_auth := models.User{}
		err := models.UserCollection.FindId(id).One(&user)
		models.UserCollection.FindId(session.Get("auth_id")).One(&user_auth)
		if err != nil {
			fmt.Println(err)
			rnd.Redirect("/")
			return
		}
		b := ""
		if id == session.Get("auth_id") {
			b = "true"
		}

		posts := []models.Post{}
		query := make(bson.M)
		query["owner"] = id
		iter := models.PostCollection.Find(query).Limit(1024).Iter()
		if err := iter.All(&posts); err == nil {
			allPost := []models.Post{}
			for i := len(posts) - 1; i >= 0; i-- {
				allPost = append(allPost, posts[i])
				fmt.Printf("%v", posts[i])
			}
			rnd.HTML(200, "user", map[string]interface{}{
				"auth_first_name": user_auth.FirstName,
				"auth_last_name":  user_auth.LastName,
				"auth_avatar":     user_auth.Avatar,
				"auth_id":         user_auth.Id,
				"first_name":      user.FirstName,
				"last_name":       user.LastName,
				"email":           user.Email,
				"avatar":          user.Avatar,
				"id":              user.Id,
				"auth_user":       b,
				"country":         user_auth.Country,
				"birthday":        user_auth.Birthday,
				"about":           user.About,
				"posts":           allPost,
			})
		}
	} else {
		rnd.HTML(200, "not_allowed", map[string]string{"error": "Not authorized"})
	}
}

func SavePost(rnd render.Render, r *http.Request, session sessions.Session) {
	if session.Get("auth_id") != "" {
		text := r.FormValue("text_post")
		title := r.FormValue("title_post")
		fmt.Println(text)
		new_post := models.Post{}

		new_post.Id = models.GenerateId()
		new_post.Owner = session.Get("auth_id").(string)
		new_post.Text = text
		new_post.Title = title
		new_post.Date = time.Now().Format(models.Layout)
		new_post.Nano = time.Now().Nanosecond()

		models.PostCollection.Insert(&new_post)
		rnd.Redirect("/user/" + session.Get("auth_id").(string))
	}
}

func AddLike(res http.ResponseWriter, r *http.Request, session sessions.Session) {
	if session.Get("auth_id") != "" {
		like_s := r.FormValue("count_like")
		//like_id := r.FormValue("id")
		like, _ := strconv.Atoi(like_s)
		like += 1

		res.Write([]byte(fmt.Sprintf(`{"counter": %d}`, like)))
	}
}

func RemovePost(res http.ResponseWriter, r *http.Request, session sessions.Session) {
	if session.Get("auth_id") != "" {
		post_id := r.FormValue("id")
		models.PostCollection.RemoveId(post_id)
		res.Write([]byte(fmt.Sprintf(`{"error": %d}`, 0)))
	}
}

func GetFollowStatus(res http.ResponseWriter, r *http.Request, session sessions.Session) {
	if session.Get("auth_id") != "" {
		my_id := session.Get("auth_id").(string)
		user_id := r.FormValue("id")
		query := make(bson.M)
		query["follower"] = my_id
		query["following"] = user_id
		count,_ := models.FollowCollection.Find(query).Count()
		status := count > 0
		res.Write([]byte(fmt.Sprintf(`{"follow_status": %v}`, status)))
	}
}

func UpdateFollowStatus(res http.ResponseWriter, r *http.Request, session sessions.Session) {
	if session.Get("auth_id") != "" {
		my_id := session.Get("auth_id").(string)
		user_id := r.FormValue("id")
		query := make(bson.M)
		query["follower"] = my_id
		query["following"] = user_id
		count,_ := models.FollowCollection.Find(query).Count()
		status := count > 0
		if status {
			models.FollowCollection.Remove(query)
		} else {
			follow_edge := models.FollowEdge{}
			follow_edge.Id = models.GenerateId()
			follow_edge.Follower = my_id
			follow_edge.Following = user_id
			models.FollowCollection.Insert(&follow_edge)
		}
		res.Write([]byte(fmt.Sprintf(`{"follow_status": %v}`, !status)))
	}
}
