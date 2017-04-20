package middleware

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
	"../"
)

func init()  {
	app.App.RegisterMiddleWare("auth",new(AuthMiddleWare))
}

type AuthMiddleWare struct {
}


func (middleware AuthMiddleWare) In(w http.ResponseWriter, r *http.Request, ps *httprouter.Params)  {
	fmt.Print("hello from auth middleware.IN")
}

func (middleware AuthMiddleWare) Out(w http.ResponseWriter, r *http.Request, ps *httprouter.Params)  {
	fmt.Print("world from auth middleware.OUT")
}
