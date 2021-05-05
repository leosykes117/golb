package env

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Specification struct {
	Port       string `envconfig:"GO_API_PORT"`
	DBHost     string `envconfig:"GOAPI_DB_HOST"`
	DBPort     string `envconfig:"GOAPI_DB_PORT"`
	DBUser     string `envconfig:"GOAPI_DB_USER"`
	DBPassword string `envconfig:"GOAPI_DB_PASSWORD"`
	DBName     string `envconfig:"GOAPI_DB_NAME"`
}

const (
	Port       = "PORT"
	DBHost     = "DB_HOST"
	DBPort     = "DB_PORT"
	DBUser     = "DB_USER"
	DBPassword = "DB_PASSWORD"
	DBName     = "DB_NAME"
)

var (
	envars Specification
)

func LoadVars() {
	filename, err := os.Executable()

	if err != nil {
		log.Fatal("Ocurrio un error al obtener la ruta de ejecución", err)
	}

	fmt.Println("filename", filename)

	dir := path.Dir(filename)

	fmt.Println("dir", dir)

	envFilePath := path.Join(path.Dir(filename), "../.env")

	fmt.Println("envFilePath", envFilePath)
	err = godotenv.Load(envFilePath)
	if err != nil {
		log.Fatal("Error al leer el archivo .env:", err)
	}
}

func ReadVars() {
	err := envconfig.Process("go_api", &envars)
	if err != nil {
		log.Fatalf("Error al leer las variables: %v", err)
	}
}

func GetEnvs(envar string) (interface{}, error) {
	switch envar {
	case Port:
		return envars.Port, nil
	case DBHost:
		return envars.DBHost, nil
	case DBPort:
		return envars.DBPort, nil
	case DBUser:
		return envars.DBUser, nil
	case DBPassword:
		return envars.DBPassword, nil
	case DBName:
		return envars.DBName, nil
	default:
		return "", fmt.Errorf("No existe la variable %s", envar)
	}
}