package app

import (
	"../core"
)


// Make sure that this file should be inited after app.go,
// that to say, the alphabet order of this file smaller than app.go file
func init()  {

	//config for app
	Config.SetByMapWithPrefix("app",map[string]string{
		"port" : "8080",
	})

	//config for database
	Config.SetByMapWithPrefix("database",map[string]string {

		"host" : core.Env("db.host","127.0.0.1"),

		"port" : core.Env("db.port","3306"),

		"database" : core.Env("db.database","dbname"),

		"user" : core.Env("db.user","root"),

		"password" : core.Env("db.password","123456"),
	})
}
