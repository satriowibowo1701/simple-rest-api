package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserController interface {
	GetAllUser(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	GetDetailUser(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}
