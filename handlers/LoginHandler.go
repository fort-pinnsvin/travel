package handlers

import (
"github.com/martini-contrib/render"
)

func LoginHandler(rnd render.Render){
	rnd.HTML(200,"signin",nil)
}
