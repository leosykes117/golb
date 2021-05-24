package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type server struct {
	*http.Server
}

func newServer(port string, router *mux.Router) *server {
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})
	return &server{
		&http.Server{
			Addr:         ":" + port,
			Handler:      handlers.CORS(headers, methods, origins)(router),
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  15 * time.Second,
		},
	}
}

func (srv *server) Start() {
	fmt.Println("Iniciando el servidor en el puerto", srv.Addr)
	srv.ListenAndServe()
}
