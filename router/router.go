package router

import (
	"github.com/julienschmidt/httprouter"
	"github.com/satriowibowo1701/simple-rest-api/controller"
)

func NewRouter(userController controller.UserController) *httprouter.Router {
	router := httprouter.New()
	router.GET("/api/users", userController.GetAllUser)
	router.GET("/api/user/:name", userController.GetDetailUser)
	return router
}
