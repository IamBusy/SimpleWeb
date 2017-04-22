package app


// Make sure that this file should be inited after app.go,
// that to say, the alphabet order of this file smaller than app.go file
func init()  {

	//config for app
	Config.SetByMapWithPrefix("app",map[string]string{
		"port" : "8080",
	})

	//config for database
	Config.SetByMapWithPrefix("database",map[string]string {

		"host" : "127.0.0.1",

		"post" : "3306",

		"database" : "name",

		"username" : "root",

		"password" : "123456",
	})
}
