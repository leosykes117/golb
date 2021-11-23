package sign

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/leosykes117/golb/backend/gocms/internal/auth"
	"github.com/leosykes117/golb/backend/gocms/internal/env"
	"github.com/leosykes117/golb/backend/gocms/pkg/models"
)

func SetAuthToken(l *models.Login, ID uint) (string, error) {
	tokenExpiration, err := env.GetEnvs(env.TokenExpiration)
	if err != nil {
		return "", err
	}

	expirtionTime, err := time.ParseDuration(tokenExpiration.(string))
	if err != nil {
		return "", err
	}

	fmt.Printf("TokenExpiration: %s\n", expirtionTime)

	claims := &models.Claims{
		ID:    ID,
		Email: l.GetEmail(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expirtionTime).Unix(),
			Issuer:    "Golb",
		},
	}
	return auth.GenerateToken(claims)
}
