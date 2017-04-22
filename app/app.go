package app

import (
	"../core"
)

var (
	App core.Container
	Router core.Router
	Config core.Config
	DB core.Manager
)

func init() {
	App = core.NewContainer()
	Config = core.NewConfig()
	Router = core.NewRouter()
	Router.SetApp(App)

	App.Put("config",Config)
	App.Put("router",Router)

	//Register DB when bootstrapping
	App.AfterBootstrap(func(container core.Container) {
		connector := core.NewConnector()
		connector.Connect(
			"mysql",
			Config.Get("database.host",""),
			Config.Get("database.port",""),
			Config.Get("database.user",""),
			Config.Get("database.password",""),
			Config.Get("database.database",""),
		)
		DB = core.NewManager()
		DB.SetConnector(connector)
		App.Put("db",DB)
		App.RegisterDB(DB)
	});

}
