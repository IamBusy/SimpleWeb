package core

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)


type Router struct {
	router *httprouter.Router

}

func (router Router)handle(method string, uri string, middlewares []MiddleWare, handler httprouter.Handle)  {

	router.router.Handle(method, uri, func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		for i:=0;i<len(middlewares);i++ {
			middlewares[i].in(w,r,ps)
		}

		handler(w,r,ps)

		for i:=len(middlewares);i>=0;i++ {
			middlewares[i].out(w,r,ps)
		}
	})
}

func (router Router) GET(uri string,middlewares []MiddleWare ,handler httprouter.Handle)  {
	router.handle("GET",uri,middlewares,handler)
}

func (router Router) POST(uri string,middlewares []MiddleWare ,handler httprouter.Handle)  {
	router.handle("POST",uri,middlewares,handler)
}

func (router Router) PUT(uri string,middlewares []MiddleWare ,handler httprouter.Handle)  {
	router.handle("PUT",uri,middlewares,handler)
}

func (router Router) DELETE(uri string,middlewares []MiddleWare ,handler httprouter.Handle)  {
	router.handle("DELETE",uri,middlewares,handler)
}