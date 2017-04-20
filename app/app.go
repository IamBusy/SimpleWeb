package app

import (
	"../core"
)

var (
	App *core.Container
	Router *core.Router
	Config *core.Config
)

func init() {
	App = &core.Container{}
	App.Init()
	Router = &core.Router{}
	Router.SetApp(App)
	Config = &core.Config{}
	Config.Init()

}
