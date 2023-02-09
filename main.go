package main

import (
	"net/http"
	"os"

	"github.com/satriowibowo1701/simple-rest-api/controller"
	"github.com/satriowibowo1701/simple-rest-api/middleware"
	"github.com/satriowibowo1701/simple-rest-api/router"
	"github.com/satriowibowo1701/simple-rest-api/service"
)

func main() {
	router := router.NewRouter(controller.NewUserController(service.NewUserService()))
	server := http.Server{
		Addr:    os.Getenv("PORT"),
		Handler: middleware.NewMiddleware(router),
	}
	server.ListenAndServe()
}
