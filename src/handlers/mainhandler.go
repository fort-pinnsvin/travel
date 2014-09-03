package handlers

import (
	"github.com/martini-contrib/render"
	//"travel/models"
)

func MainHandler(rnd render.Render){

	rnd.HTML(200, "home",nil)
}
