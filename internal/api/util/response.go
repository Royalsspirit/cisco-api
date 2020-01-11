package util

import (
	"encoding/json"
	"net/http"
)

// ResponseErrors represents the invalid request
func ResponseErrors(w http.ResponseWriter, errors []map[string]string, statusCode int) {
	e := &Errors{Errors: errors, Status: statusCode, Message: "Validation error"}

	response, _ := json.Marshal(e)

	w.WriteHeader(statusCode)
	w.Write(response)
}

// ResponseError represents the invalid request
func ResponseError(w http.ResponseWriter, message string, statusCode int) {
	e := &Error{Status: statusCode, Message: message}

	response, _ := json.Marshal(e)

	w.WriteHeader(statusCode)
	w.Write(response)
}
