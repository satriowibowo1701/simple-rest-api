package test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/satriowibowo1701/simple-rest-api/controller"
	"github.com/satriowibowo1701/simple-rest-api/middleware"
	"github.com/satriowibowo1701/simple-rest-api/router"
	"github.com/satriowibowo1701/simple-rest-api/service"
	"github.com/stretchr/testify/assert"
)

func setupRouter() http.Handler {
	router := router.NewRouter(controller.NewUserController(service.NewUserService()))

	return middleware.NewMiddleware(router)
}

func TestFindAllUserSuccess(t *testing.T) {
	router := setupRouter()
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/users", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "satrio")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"].(string))
}

func TestFindUserByNameSuccess(t *testing.T) {
	router := setupRouter()
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/user/satrio", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "satrio")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"].(string))
	assert.Equal(t, "satrio", responseBody["data"].(map[string]interface{})["login"])
}

func TestFindUserByNameNotFound(t *testing.T) {
	router := setupRouter()
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/user/-2z", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "satrio")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "Data Not Found", responseBody["status"].(string))
	assert.Equal(t, nil, responseBody["data"])
}

func TestUnauthorizhed(t *testing.T) {
	router := setupRouter()
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/users", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "test")
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, 401, response.StatusCode)
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	assert.Equal(t, 401, int(responseBody["code"].(float64)))
	assert.Equal(t, "UNAUTHORIZED", responseBody["status"].(string))
	assert.Equal(t, nil, responseBody["data"])
}
