package core

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

type MiddleWare interface {
	In(w http.ResponseWriter, r *http.Request, ps *httprouter.Params)
	Out(w http.ResponseWriter, r *http.Request, ps *httprouter.Params)
}
