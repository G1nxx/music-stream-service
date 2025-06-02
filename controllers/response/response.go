package response

import (
	"encoding/json"
	"net/http"
)


func NewErrorResponse(writer http.ResponseWriter, code int, message string) {
	writer.WriteHeader(code)
	res, _ := json.Marshal(message)
	writer.Write(res)
}

func OkResponse(writer http.ResponseWriter) {
	res, _ := json.Marshal("ok")
	writer.Write(res)
}

func LastErrorHandling(writer http.ResponseWriter, err error) {
	switch err.(type) {
	// case *domainErrors.InvalidField:
	// 	NewErrorResponse(writer, http.StatusBadRequest, err.Error())
	// 	return
	default:
		NewErrorResponse(writer, http.StatusInternalServerError, err.Error())
		return
	}
}

func InfoResponse(writer http.ResponseWriter, code int, body []byte) {
	writer.WriteHeader(code)
	writer.Write(body)
}