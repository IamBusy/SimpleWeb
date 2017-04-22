package app

import (
	"../core"
)

var (
	App core.Container
	Router core.Router
	Config core.Config
	DB core.DBManager
)

func init() {
	App = core.NewContainer()
	Config = core.NewConfig()
	Router = core.NewRouter()
	Router.SetApp(App)

	//DB = &core.DBManager{}
	//connector := &core.DBConnector{}
	//connector.Connect("mysql",
	//	Config.Get("database.host",""),
	//	Config.Get("database.port",""),
	//	Config.Get("database.user",""),
	//	Config.Get("database.password",""),
	//	Config.Get("database.database",""))
	//
	//fmt.Print(Config.Get("database.host","123"))
	//
	//DB.SetConnector(connector)

}
