package main

import (
	_ "./app/middleware"
	_ "./app/controller"
	"net/http"
	"./app"
	"log"

)


func main()  {
	app.App.Bootstrap()
	log.Fatal(http.ListenAndServe(":"+app.Config.Get("app.port","8080"), app.Router))
}
