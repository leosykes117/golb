package auth

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
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
