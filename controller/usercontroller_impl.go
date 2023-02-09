package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/satriowibowo1701/simple-rest-api/helper"
	"github.com/satriowibowo1701/simple-rest-api/service"
)

type UserControllerImpl struct {
	UserService service.UserServiceImpl
}

func NewUserController(user service.UserServiceImpl) UserController {
	return &UserControllerImpl{user}
}

func (user *UserControllerImpl) GetAllUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	data, err := user.UserService.GetUsers()
	response := helper.ResponseWithData(err, data)
	helper.WriteToResponseBody(w, response, err, http.StatusBadRequest)
}

func (user *UserControllerImpl) GetDetailUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	name := params.ByName("name")
	data, err := user.UserService.GetDetailUser(name)
	response := helper.ResponseWithData(err, data)
	helper.WriteToResponseBody(w, response, err, http.StatusBadRequest)
}
