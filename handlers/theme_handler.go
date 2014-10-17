package handlers

import (
	"fmt"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"net/http"
	"os"
)

func LoadTheme(res http.ResponseWriter, session sessions.Session) {
	fi, _ := os.Open("assets/css/bootstrap.css")
	if session.Get("theme") == "White" {
		fi, _ = os.Open("assets/css/w_bootstrap.mi.css")
	} else {
		fi, _ = os.Open("assets/css/b_bootstrap.mi.css")
	}

	stat, _ := fi.Stat()

	buf := make([]byte, stat.Size())
	fi.Read(buf)
	res.Header().Set("Content-Type", "text/css")
	res.Write([]byte(buf))

}

func SetTheme(rnd render.Render, res http.ResponseWriter, r *http.Request, session sessions.Session) {
	theme := r.FormValue("theme")
	fmt.Println(theme)
	session.Set("theme", theme)
	//LoadTheme(res,session)
	rnd.Redirect("/")
}
