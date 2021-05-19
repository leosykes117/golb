package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/leosykes117/golb/backend/gocms/pkg/models"
	"github.com/leosykes117/golb/backend/gocms/pkg/models/user"
	"github.com/leosykes117/golb/backend/gocms/pkg/sign"
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
		log.Fatalf("ERROR MySQL on user.Create: %v", err)
	}

	lg := models.NewLogIn(user.Email, user.PasswordHash)
	tokenStr, err := sign.SetAuthToken(lg, user.ID)

	if err != nil {
		ErrCreateToken.Send(w)
	}

	log.Printf("Usuario %v creado con exito", user)

	data := make(map[string]string)
	data["name"] = user.FullName()
	data["gender"] = user.Gender
	data["token"] = tokenStr

	Success(data, http.StatusCreated).Send(w)
}

func (s *Services) loginHandler(w http.ResponseWriter, r *http.Request) {
	data := new(models.Login)
	body := r.Body
	defer body.Close()

	if err := data.Unmarshal(body); err != nil {
		log.Printf("ERROR loginHandler. No se pudo leer el payload: %v", err)
		ErrInvalidJSON.Send(w)
		return
	}

	log.Printf("Datos de login: %v", data)

	user, err := s.userService.LogIn(data)
	if err != nil {
		log.Fatalf("ERROR MySQL on user.LogIn %v", err)
	}

	lg := models.NewLogIn(user.Email, user.PasswordHash)
	tokenStr, err := sign.SetAuthToken(lg, user.ID)

	if err != nil {
		ErrCreateToken.Send(w)
		return
	}

	log.Printf("Usuario %v creado con exito", user)

	responseData := make(map[string]string)
	responseData["name"] = user.FullName()
	responseData["gender"] = user.Gender
	responseData["token"] = tokenStr

	Success(responseData, http.StatusCreated).Send(w)
}
