package utils

import (
	"go-risko/data"

	"encoding/json"
	"net/http"
)

func JsonSuccessResponse(w http.ResponseWriter, Data interface{}, msg string) {
	result, _ := json.Marshal(data.Response{
		Code: http.StatusOK,
		Data: Data,
		Message: msg,
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func JsonErrorResponse(w http.ResponseWriter, statusCode int, msg string) {
	result, _ := json.Marshal(data.Response{
		Code: statusCode,
		Message: msg,
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(result)
}