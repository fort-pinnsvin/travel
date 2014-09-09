package handlers

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"labix.org/v2/mgo/bson"
	"models"
	"net/http"
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
				"about":           user_auth.About,
				"posts":           allPost,
			})
		}
	} else {
		rnd.HTML(200, "not_allowed", map[string]string{"error": "Not authorized"})
	}
}

func SavePost(res http.ResponseWriter, rnd render.Render, r *http.Request, session sessions.Session) {
	if session.Get("auth_id") != "" {
		text := r.FormValue("text_post")
		title := r.FormValue("title_post")
		fmt.Println(title)
		new_post := models.Post{}

		new_post.Id = models.GenerateId()
		new_post.Owner = session.Get("auth_id").(string)
		new_post.Text = text
		new_post.Title = title


		models.PostCollection.Insert(&new_post)
		res.Write([]byte(fmt.Sprintf(`{"id_user": %s}`, session.Get("auth_id").(string))))
	}
}

func AddLike(res http.ResponseWriter, r *http.Request, session sessions.Session) {
	like_id := r.FormValue("id")
	q := make(bson.M)
	q["liker"] = session.Get("auth_id").(string)
	q["idpost"] = like_id
	like_arr := models.LikeCollection.Find(q).Limit(10).Iter()
	likes := []models.Like{}
	_ = like_arr.All(&likes)
	fmt.Println(likes)
	status := false
	if session.Get("auth_id") != "" {
		post := models.Post{}
		models.PostCollection.FindId(like_id).One(&post)
		if len(likes) == 0 {
			like := models.Like{session.Get("auth_id").(string), like_id}
			models.LikeCollection.Insert(&like)
			post.Like += 1
			status = true
		}else{
			models.LikeCollection.Remove(q)
			post.Like -= 1
			status = false
		}
		models.PostCollection.UpdateId(like_id,post)

		res.Write([]byte(fmt.Sprintf(`{"counter": %d,"status_like": %v}`, post.Like, status)))
	}
}

