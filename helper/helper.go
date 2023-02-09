package helper

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/satriowibowo1701/simple-rest-api/model"
)

func ApiCall(method, url string, body io.Reader, data interface{}) error {
	var client = &http.Client{}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return err
	}
	res, errclient := client.Do(req)
	if errclient != nil {
		return errclient

	}
	if res.StatusCode == 404 {
		return errors.New("Data Not Found")
	}
	defer res.Body.Close()
	json.NewDecoder(res.Body).Decode(data)
	return nil
}

func ResponseWithData(err error, data interface{}) interface{} {

	if err != nil {
		var status string
		var code int

		if err.Error() == "Data Not Found" {
			status = err.Error()
			code = http.StatusNotFound
		} else {
			status = "Bad Request"
			code = http.StatusBadRequest
		}
		response := model.WebResponse{
			Code:   code,
			Status: status,
			Data:   nil,
		}
		return response

	}

	response := model.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   data,
	}
	return response

}

func WriteToResponseBody(writer http.ResponseWriter, response interface{}, err error, httpcode int) {
	writer.Header().Add("Content-Type", "application/json")
	if err != nil {
		var code = httpcode
		if err.Error() == "Data Not Found" {
			code = http.StatusNotFound
		}
		writer.WriteHeader(code)
	}
	encoder := json.NewEncoder(writer)
	_ = encoder.Encode(response)
}
