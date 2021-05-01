package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func routes(services *Services) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/users", services.createUserHandler).Methods(http.MethodPost)
	//r.HandleFunc("/api/gophers", s.FetchGophers).Methods(http.MethodGet)
	//r.HandleFunc("/api/gophers/{ID:[a-zA-Z0-9_]+}", s.FetchGopher).Methods(http.MethodGet)
	//r.HandleFunc("/api/gophers/{ID:[a-zA-Z0-9_]+}", s.ModifyGopher).Methods(http.MethodPut)
	//r.HandleFunc("/api/gophers/{ID:[a-zA-Z0-9_]+}", s.RemoveGopher).Methods(http.MethodDelete)

	return r
}
