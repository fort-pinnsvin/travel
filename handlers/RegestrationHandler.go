package handlers

import (
	"github.com/martini-contrib/render"
)

func RegestrationHandler(rnd render.Render){
	rnd.HTML(200,"signup",nil)
}
