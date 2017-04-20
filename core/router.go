package core

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)


/**
 * Router
 */
type Router struct {
	app *Container
	router *httprouter.Router
}

func (router *Router) SetApp(container *Container)  {
	router.app = container
}

/**
 * the entry point for applying middleware and handling request
 */
func (router *Router)handle(method string, uri string, middlewareNames []string, handler httprouter.Handle)  {


	if router.router == nil {
		router.router = httprouter.New()
	}

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

func (router *Router) GET(uri string,middlewareNames []string ,handler httprouter.Handle)  {
	router.handle("GET",uri,middlewareNames,handler)
}

func (router *Router) POST(uri string,middlewareNames []string ,handler httprouter.Handle)  {
	router.handle("POST",uri,middlewareNames,handler)
}

func (router *Router) PUT(uri string,middlewareNames []string ,handler httprouter.Handle)  {
	router.handle("PUT",uri,middlewareNames,handler)
}

func (router *Router) DELETE(uri string,middlewareNames []string ,handler httprouter.Handle)  {
	router.handle("DELETE",uri,middlewareNames,handler)
}

func (router *Router) ServeHTTP(w http.ResponseWriter, req *http.Request)  {
	router.router.ServeHTTP(w,req)
}