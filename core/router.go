package core

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"fmt"
)


/**
 * router
 */
type Router interface {
	SetApp(container Container)
	GET(uri string,middlewareNames []string ,handler Handler)
	POST(uri string,middlewareNames []string ,handler Handler)
	PUT(uri string,middlewareNames []string ,handler Handler)
	DELETE(uri string,middlewareNames []string ,handler Handler)
	ServeHTTP(w http.ResponseWriter, req *http.Request)
}

type Handler func(*http.Request, httprouter.Params) string

type router struct {
	app Container
	router *httprouter.Router
}

func NewRouter() Router {
	rutr := &router{}
	rutr.router = httprouter.New()
	return rutr
}

func (router *router) SetApp(container Container)  {
	router.app = container
}

/**
 * the entry point for applying middleware and handling request
 */
func (router *router)handle(method string, uri string, middlewareNames []string, handler Handler)  {


	router.router.Handle(method, uri, func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		middlewares := make([]*MiddleWare,0,5) ;

		if middlewareNames != nil {
			for i:=0; i<len(middlewareNames);i++  {
				middleware := router.app.Middleware(middlewareNames[i])
				if middleware != nil {
					middlewares = append(middlewares,middleware)
				}
			}
		}

		for i:=0;i<len(middlewares);i++ {
			if middlewares[i] != nil {
				(*middlewares[i]).In(w,r,&ps)
			}

		}

		output := handler(r,ps)
		fmt.Fprint(w, output)

		for i:=len(middlewares)-1;i>=0;i-- {
			if middlewares[i] != nil {
				(*middlewares[i]).Out(w,r,&ps)
			}
		}
	})
}

func (router *router) GET(uri string,middlewareNames []string ,handler Handler)  {
	router.handle("GET",uri,middlewareNames,handler)
}

func (router *router) POST(uri string,middlewareNames []string ,handler Handler)  {
	router.handle("POST",uri,middlewareNames,handler)
}

func (router *router) PUT(uri string,middlewareNames []string ,handler Handler)  {
	router.handle("PUT",uri,middlewareNames,handler)
}

func (router *router) DELETE(uri string,middlewareNames []string ,handler Handler)  {
	router.handle("DELETE",uri,middlewareNames,handler)
}

func (router *router) ServeHTTP(w http.ResponseWriter, req *http.Request)  {
	router.router.ServeHTTP(w,req)
}