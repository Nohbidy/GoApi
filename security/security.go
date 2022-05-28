package security

import (
	configs "fiber-mongo-api/config"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func CreateJWTToken() (string, int64, error) {
	exp := time.Now().Add(time.Minute * 15).Unix()

	claims := jwt.MapClaims{"exp": exp}
	unsigned := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := configs.EnvJWTSecret()
	signedToken, err := unsigned.SignedString([]byte(secret))
	if err != nil {
		return "", 0, err
	}

	return signedToken, exp, nil
}
