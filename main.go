package main

import (
	_ "./app/middleware"
	"net/http"
	"./app"
	"log"
	_ "./app/controller"
)


func main()  {
	app.App.Bootstrap()
	log.Fatal(http.ListenAndServe(":"+app.Config.Get("app.port","8080"), app.Router))
}
