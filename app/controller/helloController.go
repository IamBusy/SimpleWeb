package controller

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
	"../"
	"../service"
)

var (
	Hello *helloController
)

func init() {
	Hello = &helloController{}
	app.Router.GET("/hello",[]string{"auth"},Hello.GET)
}

type helloController struct {

}

func (controller helloController) GET(r *http.Request, ps httprouter.Params) string  {
	result := service.UserService.All()
	fmt.Print(result.Columns())
	return "hello from helloController"
}