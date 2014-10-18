package handlers

import (
	"fmt"
	"github.com/fort-pinnsvin/travel/helpfunc"
	"github.com/fort-pinnsvin/travel/models"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"labix.org/v2/mgo/bson"
	"net/http"
	"sort"
	"time"
	"github.com/fort-pinnsvin/travel/utils"
)

func MiniBlogHandler(rnd render.Render, session sessions.Session, params martini.Params) {
	if session.Get("auth_id") != "" {
		id_blog := params["id"]

		postsBlog := []models.PostBlog{}
		query := make(bson.M)
		query["idblog"] = id_blog
		models.PostBlogCollection.Find(query).Limit(1024).Iter().All(&postsBlog)

		blog := models.Blog{}
		models.BlogCollection.FindId(id_blog).One(&blog)

		sort.Sort(models.ByPostBlog(postsBlog))

		user_auth := helpfunc.GetAuthUser(session)
		rnd.HTML(200, "blog", map[string]interface{}{
			"user":      user_auth,
			"auth_user": true,
			"id_blog":   id_blog,
			"all_post":  postsBlog,
			"blog":      blog,
		})
	}
}

func MiniBlogListHandler(rnd render.Render, session sessions.Session) {
	if session.Get("auth_id") != "" {
		id_user := session.Get("auth_id").(string)
		blogs := []models.Blog{}
		query := make(bson.M)
		query["owner"] = id_user
		models.BlogCollection.Find(query).Limit(1024).Iter().All(&blogs)

		sort.Sort(models.ByBlog(blogs))

		user_auth := helpfunc.GetAuthUser(session)
		rnd.HTML(200, "lists_blogs", map[string]interface{}{
			"user":      user_auth,
			"auth_user": true,
			"blogs":     blogs,
		})
	}
}

func SavePostMiniblog(res http.ResponseWriter, rnd render.Render, r *http.Request, session sessions.Session) {
	if session.Get("auth_id") != "" {
		text := r.FormValue("text_post")
		id_blog := r.FormValue("id_blog")

		new_post := models.PostBlog{}

		new_post.Id = models.GenerateId()
		new_post.IdBlog = id_blog
		new_post.Owner = session.Get("auth_id").(string)
		new_post.Text = text
		new_post.Date = time.Now().Format(models.Layout)
		new_post.Nano = time.Now().Unix()

		models.PostBlogCollection.Insert(&new_post)
		res.Write([]byte(fmt.Sprintf(`{"id_blog": "%s"}`, id_blog)))
	}
}

func CreateMiniBlog(res http.ResponseWriter, rnd render.Render, r *http.Request, session sessions.Session) {
	if session.Get("auth_id") != "" {
		new_blog := models.Blog{}

		new_blog.Id = models.GenerateId()
		new_blog.Name = "New Story"
		new_blog.Owner = session.Get("auth_id").(string)
		new_blog.Date = time.Now().Format(models.Layout)
		new_blog.Nano = time.Now().Unix()

		new_post := models.Post{}

		new_post.Id = models.GenerateId()
		new_post.Owner = session.Get("auth_id").(string)
		new_post.Text = `Read it <a href="` +
				"//" + utils.GetValue("WWW", "localhost:3000") + "/mini_blog/" + new_blog.Id +`">here</a> .`
		new_post.Title = "Oh, watch my new Trip!"
		new_post.Date = time.Now().Format(models.Layout)
		new_post.Nano = time.Now().Unix()

		models.PostCollection.Insert(&new_post)
		models.BlogCollection.Insert(&new_blog)
		res.Write([]byte(fmt.Sprintf(`{"id_blog": "%s"}`, new_blog.Id)))
	}
}

func RemovePostMiniblog(res http.ResponseWriter, r *http.Request, session sessions.Session) {
	if session.Get("auth_id") != "" {
		post_id := r.FormValue("id")
		models.PostBlogCollection.RemoveId(post_id)
		res.Write([]byte(fmt.Sprintf(`{"St": "%s"}`, "OK" )))
	}
}

func MiniBlogEdit(rnd render.Render, session sessions.Session, params martini.Params){
	if session.Get("auth_id") != "" {
		user_auth := helpfunc.GetAuthUser(session)
		id_blog := params["id"]
		blog := models.Blog{}
		models.BlogCollection.FindId(&id_blog).One(&blog)

		rnd.HTML(200, "edit_blog", map[string]interface {}{
			"user" : user_auth,
			"auth_user" : true,
			"blog" : blog,
		});
	}
}


func SaveEditBlog(res http.ResponseWriter, session sessions.Session, r *http.Request,  params martini.Params) {
	if session.Get("auth_id") != "" {
		id := params["id"]
		blog := &models.Blog{}
		models.BlogCollection.FindId(id).One(&blog)
		blog.Name = r.FormValue("title")
		models.BlogCollection.UpdateId(id, blog)
		res.Write([]byte(`{"error": 0}`))
	} else {
		res.Write([]byte(`{"error": 1}`))
	}
}
