package core

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)


/**
 * router
 */
type Router interface {
	SetApp(container Container)
	GET(uri string,middlewareNames []string ,handler httprouter.Handle)
	POST(uri string,middlewareNames []string ,handler httprouter.Handle)
	PUT(uri string,middlewareNames []string ,handler httprouter.Handle)
	DELETE(uri string,middlewareNames []string ,handler httprouter.Handle)
	ServeHTTP(w http.ResponseWriter, req *http.Request)
}

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
func (router *router)handle(method string, uri string, middlewareNames []string, handler httprouter.Handle)  {


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

		handler(w,r,ps)

		for i:=len(middlewares)-1;i>=0;i-- {
			if middlewares[i] != nil {
				(*middlewares[i]).Out(w,r,&ps)
			}
		}
	})
}

func (router *router) GET(uri string,middlewareNames []string ,handler httprouter.Handle)  {
	router.handle("GET",uri,middlewareNames,handler)
}

func (router *router) POST(uri string,middlewareNames []string ,handler httprouter.Handle)  {
	router.handle("POST",uri,middlewareNames,handler)
}

func (router *router) PUT(uri string,middlewareNames []string ,handler httprouter.Handle)  {
	router.handle("PUT",uri,middlewareNames,handler)
}

func (router *router) DELETE(uri string,middlewareNames []string ,handler httprouter.Handle)  {
	router.handle("DELETE",uri,middlewareNames,handler)
}

func (router *router) ServeHTTP(w http.ResponseWriter, req *http.Request)  {
	router.router.ServeHTTP(w,req)
}