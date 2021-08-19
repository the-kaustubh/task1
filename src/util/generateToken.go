package util

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type CustomClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func SignToken(username string) (string, error) {
	claims := CustomClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).UTC().Unix(),
			Issuer:    "GolangGinTrial",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("secretsecretsecretsecretsecretsecret"))
	if err != nil {
		return "", err
	}
	return signedToken, nil

}
