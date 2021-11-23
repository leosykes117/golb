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
		log.Printf("ERROR createUserHandler. No se pudo leer el json: %v\n", err)
		ErrInvalidJSON.Send(w)
		return
	}

	log.Printf("Datos del usuario: %v\n", user)

	if err := s.userService.Create(user); err != nil {
		log.Printf("ERROR MySQL on user.Create: %v\n", err)
		ErrBadRequest.Send(w)
		return
	}

	lg := models.NewLogIn(user.Email, user.PasswordHash)
	tokenStr, err := sign.SetAuthToken(lg, user.ID)

	if err != nil {
		ErrCreateToken.Send(w)
		return
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
		log.Printf("ERROR MySQL on user.LogIn %v", err)
		ErrUserNotFound.Send(w)
		return
	}

	lg := models.NewLogIn(user.Email, user.PasswordHash)
	tokenStr, err := sign.SetAuthToken(lg, user.ID)

	if err != nil {
		ErrCreateToken.Send(w)
		return
	}

	log.Printf("Usuario %v autenticado con éxito", user)

	responseData := make(map[string]string)
	responseData["name"] = user.FullName()
	responseData["gender"] = user.Gender
	responseData["token"] = tokenStr

	Success(responseData, http.StatusCreated).Send(w)
}

func (s *Services) testAuthorizationHandler(w http.ResponseWriter, r *http.Request) {
	body := r.Body
	defer body.Close()

	reqBody := make(map[string]string)
	if err := json.NewDecoder(body).Decode(&reqBody); err != nil {
		log.Printf("ERROR testAuthorizationHandler. No se pudo obtener el cuerpo de la petición: %v\n", err)
		ErrInvalidJSON.Send(w)
		return
	}

	log.Printf("Datos de la peticón: %v\n", reqBody)

	data := make(map[string]string)
	data["message"] = "Autenticado para realizar peticiones."
	Success(data, http.StatusOK).Send(w)
}
