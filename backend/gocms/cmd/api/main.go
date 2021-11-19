package main

import (
	"log"
	"os"

	"github.com/leosykes117/golb/backend/gocms/internal/auth"
	"github.com/leosykes117/golb/backend/gocms/internal/env"
	"github.com/leosykes117/golb/backend/gocms/pkg/api"
)

func main() {
	fromDocker, ok := os.LookupEnv("FROM_DOCKER")
	log.Printf("fromDocker: %s\nok: %t\n\n", fromDocker, ok)

	if !ok {
		env.LoadVars()
	}

	env.ReadVars()

	privateFile, err := env.GetEnvs(env.PrivateKeyPath)
	if err != nil {
		log.Fatal("No existe el archivo de la llave privada")
	}

	publicFile, err := env.GetEnvs(env.PublicKeyPath)
	if err != nil {
		log.Fatal("No existe el archivo de la llave pública")
	}

	err = auth.LoadKeys(privateFile.(string), publicFile.(string))
	if err != nil {
		log.Fatalf("ERROR LoadKeys() %v", err)
	}

	serverPort, err := env.GetEnvs(env.Port)
	if err != nil {
		log.Fatal("Ocurrio un error al obtener la variable Port")
	}

	log.Printf("GO_API_PORT %s", serverPort)

	api.Start(serverPort.(string))
}
