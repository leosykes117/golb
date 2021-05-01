package main

import (
	"log"
	"os"

	"github.com/leosykes117/golb/backend/gocms/internal/env"
	"github.com/leosykes117/golb/backend/gocms/pkg/api"
)

func main() {
	_, ok := os.LookupEnv("FROM_DOCKER")

	if !ok {
		env.LoadVars()
	}

	env.ReadVars()

	serverPort, err := env.GetEnvs(env.Port)
	if err != nil {
		log.Fatal("Ocurrio un error al obtener la variable Port")
	}

	log.Printf("GO_API_PORT %s", serverPort)

	api.Start(serverPort.(string))
}
