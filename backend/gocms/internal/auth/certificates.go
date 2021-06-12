package auth

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"sync"

	"github.com/dgrijalva/jwt-go"
)

var (
	signKey   *rsa.PrivateKey
	verifyKey *rsa.PublicKey
	once      sync.Once
)

// LoadKeys .
func LoadKeys(privateFile, publicFile string) error {
	var err error
	once.Do(func() {
		err = loadFiles(privateFile, publicFile)
	})
	return err
}

func loadFiles(privateFile, publicFile string) error {
	fmt.Println(privateFile, "--", publicFile)
	privateFile = setPath(privateFile)
	publicFile = setPath(publicFile)
	fmt.Println(privateFile, "--", publicFile)
	privKeyBytes, err := ioutil.ReadFile(privateFile)
	if err != nil {
		return fmt.Errorf("Ocurrio un error al leer el archivo de la llave privada: %v", err)
	}

	pubKeyBytes, err := ioutil.ReadFile(publicFile)
	if err != nil {
		return fmt.Errorf("Ocurrio un error al leer el archivo de la llave pública: %v", err)
	}

	return parseRSA(privKeyBytes, pubKeyBytes)
}

func parseRSA(privateBytes, publicBytes []byte) error {
	var err error
	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		return fmt.Errorf("Ocurrio un error al parsear la llave privada: %v", err)
	}
	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		return fmt.Errorf("Ocurrio un error al parsear la llave píblica: %v", err)
	}
	return nil
}

func setPath(file string) string {
	if filepath.IsAbs(file) {
		return file
	}
	filename, err := os.Executable()

	if err != nil {
		log.Fatal("Ocurrio un error al obtener la ruta de ejecución", err)
	}

	fmt.Println("filename", filename)

	dir := path.Dir(filename)

	abs, _ := filepath.Abs(filepath.Join(dir, file))
	return abs
}
