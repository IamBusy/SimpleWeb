package core

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

type MiddleWare interface {
	in(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	out(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}
