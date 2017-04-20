package controller

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
)

var (
	Hello *helloController
)

func init() {
	Hello = &helloController{}
}

type helloController struct {

}

func (controller helloController) GET(w http.ResponseWriter, r *http.Request, ps httprouter.Params)  {
	fmt.Fprint(w, "Hello World from HelloController")
}