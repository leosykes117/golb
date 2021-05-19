package sign

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/leosykes117/golb/backend/gocms/internal/auth"
	"github.com/leosykes117/golb/backend/gocms/pkg/models"
)

func SetAuthToken(l *models.Login, ID uint) (string, error) {
	claims := &models.Claims{
		ID:    ID,
		Email: l.GetEmail(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			Issuer:    "Golb",
		},
	}
	return auth.GenerateToken(claims)
}
