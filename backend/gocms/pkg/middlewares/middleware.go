package middlewares

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/leosykes117/golb/backend/gocms/internal/auth"
)

type middlewareResponse map[string]interface{}

var (
	ErrInvalidToken = middlewareResponse{"statusCode": http.StatusForbidden, "type": "toke_invalid", "message": "El token es invalido"}
)

// Authentication es el middleware para poder validar el toquen del usuario
func Authentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		_, err := auth.ValidateToken(token)
		if err != nil {
			log.Printf("ERROR Middlewares.Authentication %v", err)
			forbidden(w)
			return
		}
		next(w, r)
	}
}

func forbidden(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusForbidden)
	err := json.NewEncoder(w).Encode(ErrInvalidToken)

	if err != nil {
		log.Printf("Error al convertir la estructura usuario en json: %v", err)
	}
}
