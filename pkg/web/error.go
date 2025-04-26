package web

import (
	"encoding/json"
	"errors"
	"net/http"
)

// todo refactor this file and create a better error handling
// todo: all the web package must be improved

var (
	ErrNotFound      = errors.New("not found")       // 404
	ErrInternal      = errors.New("internal error")  // 500
	ErrBadRequest    = errors.New("invalid request") // 400
	ErrFieldNotFound = errors.New("field not found") // 422
	ErrConflict      = errors.New("conflict")        // 409
)

type errorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func Error(w http.ResponseWriter, statusCode int, message string) {
	// default status code
	defaultStatusCode := http.StatusInternalServerError
	// check if status code is valid
	if statusCode > 299 && statusCode < 600 {
		defaultStatusCode = statusCode
	}

	// response
	body := errorResponse{
		Status:  http.StatusText(defaultStatusCode),
		Message: message,
	}
	bytes, err := json.Marshal(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// write response
	w.WriteHeader(defaultStatusCode)
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}
