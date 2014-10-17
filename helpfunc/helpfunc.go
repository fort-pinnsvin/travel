package helpfunc

import (
	"github.com/fort-pinnsvin/travel/models"
	"github.com/martini-contrib/sessions"
)

func GetAuthUser(session sessions.Session) models.User {
	user := models.User{}
	user.FirstName = session.Get("first_name").(string)
	user.LastName = session.Get("last_name").(string)
	user.Id = session.Get("auth_id").(string)
	user.Avatar = session.Get("avatar").(string)
	return user
}
