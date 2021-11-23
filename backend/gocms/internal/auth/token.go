package auth

import (
	"errors"
	"log"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(claim jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claim)
	return token.SignedString(signKey)
}

func ValidateToken(t string) (jwt.MapClaims, error) {
	token, err := jwt.ParseWithClaims(t, &jwt.MapClaims{}, getVerifyKey)
	if err != nil {
		v, _ := err.(*jwt.ValidationError)
		log.Printf("jwt.ValidationError: %v", v)
		if v.Errors == jwt.ValidationErrorExpired {
			err = errors.New("token_expired")
		}
		return jwt.MapClaims{}, err
	}
	if !token.Valid {
		return jwt.MapClaims{}, errors.New("Token erroneo")
	}
	claim, ok := token.Claims.(*jwt.MapClaims)
	log.Printf("token.Claims: %v => %T", token.Claims, token.Claims)
	if !ok {
		return jwt.MapClaims{}, errors.New("Los claims no pudieron ser obtenidos")
	}
	return *claim, nil
}

func getVerifyKey(t *jwt.Token) (interface{}, error) {
	return verifyKey, nil
}
