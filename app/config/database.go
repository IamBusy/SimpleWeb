package config

import (
	"../"
)


func init()  {
	app.Config.SetByMapWithPrefix("database",map[string]string {

		"host" : "127.0.0.1",

		"post" : "3306",

		"database" : "name",

		"username" : "root",

		"password" : "123456",
	})
}