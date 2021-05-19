package api

import (
	"encoding/json"
	"net/http"
)

var (
	ErrBadRequest   = &UserError{StatusCode: http.StatusBadRequest, Type: "api_error", Message: "Cannot process current request"}
	ErrUserNotFound = &UserError{StatusCode: http.StatusBadRequest, Type: "user_not_found", Message: "Cannot find the given username"}
	ErrInvalidJSON  = &UserError{StatusCode: http.StatusBadRequest, Type: "invalid_json", Message: "Invalid or malformed JSON"}
	ErrCreateToken  = &UserError{StatusCode: http.StatusInternalServerError, Type: "token_marshal_error", Message: "Cannot create authorization"}
)

type UserError struct {
	StatusCode int    `json:"code"`
	Type       string `json:"type"`
	Message    string `json:"message,omitempty"`
}

func (e *UserError) Send(w http.ResponseWriter) error {
	statusCode := e.StatusCode
	if statusCode == 0 {
		statusCode = http.StatusBadRequest
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(e)
}
