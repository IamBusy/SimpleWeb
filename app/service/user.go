package service

import (
	"../../core"
	"database/sql"
	"../"
)

var(
	UserService *user
)

func init()  {
	UserService = &user{}
	UserService.SetApp(app.App)
}

type user struct {
	app core.Container
}

func (user *user)SetApp(container core.Container)  {
	user.app = container
}

func (user *user) All() *sql.Rows {
	result,_ := user.app.DB().Raw("select * from users")
	return result
}
