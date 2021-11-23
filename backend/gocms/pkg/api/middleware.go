package api

import (
	"log"
	"net/http"
	"strings"

	"github.com/leosykes117/golb/backend/gocms/internal/auth"
)

type middlewareResponse map[string]interface{}

var (
	ErrInvalidToken = middlewareResponse{"statusCode": http.StatusUnauthorized, "type": "toke_invalid", "message": "El token es invalido"}
)

// AuthenticationMiddleware es el middleware para poder validar el toquen del usuario
func AuthenticationMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		log.Printf("Header[\"Authorization\"]: %q, validation: %t", token, token == "")
		if token == "" {
			ErrUnauthorized.Send(w)
			return
		}
		token = strings.Split(token, "Bearer ")[1]
		log.Printf("token: %q", token)
		claims, err := auth.ValidateToken(token)
		log.Printf("Claims: %s", claims)
		if err != nil {
			log.Printf("ERROR Middlewares.Authentication %v", err)
			authError := &UserError{}

			switch err.Error() {
			case ErrTokenExpired.Type:
				authError = ErrTokenExpired
			default:
				authError = ErrUnauthorized
			}

			authError.Send(w)
			return
		}
		next(w, r)
	}
}
