package middleware

import (
	"errors"
	"net/http"

	"github.com/satriowibowo1701/simple-rest-api/helper"
	"github.com/satriowibowo1701/simple-rest-api/model"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("X-API-Key") == "satrio" {
		middleware.Handler.ServeHTTP(w, r)
	} else {
		w.Header().Set("Content-Type", "application/json")
		Response := model.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
			Data:   nil,
		}
		helper.WriteToResponseBody(w, Response, errors.New("UNAUTHORIZED"), http.StatusUnauthorized)
	}
}
