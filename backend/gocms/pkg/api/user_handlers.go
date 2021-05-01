package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/leosykes117/golb/backend/gocms/pkg/models/user"
)

func (s *Services) createUserHandler(w http.ResponseWriter, r *http.Request) {
	user := new(user.Model)
	body := r.Body
	defer body.Close()

	log.Println("Creando un nuevo usuario")

	if err := json.NewDecoder(body).Decode(user); err != nil {
		log.Printf("ERROR createUserHandler. No se pudo leer el json: %v", err)
		ErrInvalidJSON.Send(w)
		return
	}

	log.Printf("Datos del usuario: %v", user)

	if err := s.userService.Create(user); err != nil {
		log.Fatalf("ERROR MySQL on product.Create: %v", err)
	}

	log.Printf("Usuario %v creado con exito", user)

	Success(user, http.StatusCreated).Send(w)
}
