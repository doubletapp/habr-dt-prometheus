package api

import (
	"encoding/json"
	"net/http"
)

func HandleJson(response http.ResponseWriter, status int, responseBody []byte) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(status)
	response.Write(responseBody)
	return
}

func HandleError(response http.ResponseWriter, status int, message string) {
	responseBody, _ := json.Marshal(DeafultResponse{Success: false, Message: message})
	HandleJson(response, status, responseBody)
	return
}

func HandleSuccess(response http.ResponseWriter) {
	responseBody, _ := json.Marshal(DeafultResponse{Success: true})
	HandleJson(response, http.StatusOK, responseBody)
	return
}
