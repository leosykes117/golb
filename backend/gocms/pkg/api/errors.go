package api

import (
	"encoding/json"
	"net/http"
)

var (
	ErrBadRequest   = &UserError{StatusCode: http.StatusBadRequest, Type: "api_error", Message: "Cannot process current request"}
	ErrUserNotFound = &UserError{StatusCode: http.StatusBadRequest, Type: "user_not_found", Message: "Correo electrónico o contraseña incorrecta"}
	ErrInvalidJSON  = &UserError{StatusCode: http.StatusBadRequest, Type: "invalid_json", Message: "Invalid or malformed JSON"}
	ErrCreateToken  = &UserError{StatusCode: http.StatusInternalServerError, Type: "token_marshal_error", Message: "Cannot create authorization"}
	ErrUnauthorized = &UserError{StatusCode: http.StatusUnauthorized, Type: "invalid_token", Message: "El token de autorización no es válido"}
	ErrTokenExpired = &UserError{StatusCode: http.StatusUnauthorized, Type: "token_expired", Message: "El token ha expirado"}
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
