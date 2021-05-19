package auth

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(claim jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claim)
	return token.SignedString(signKey)
}

func ValidateToken(t string) (jwt.MapClaims, error) {
	token, err := jwt.ParseWithClaims(t, &jwt.MapClaims{}, getVerifyKey)
	if err != nil {
		return jwt.MapClaims{}, err
	}
	if !token.Valid {
		return jwt.MapClaims{}, errors.New("Token erroneo")
	}
	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return jwt.MapClaims{}, errors.New("Los claims no pudieron ser obtenidos")
	}
	return claim, nil
}

func getVerifyKey(t *jwt.Token) (interface{}, error) {
	return verifyKey, nil
}
