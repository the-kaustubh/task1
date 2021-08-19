package middlewares

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/the-kaustubh/task1/util"
)

func VerifyToken(c *gin.Context) {
	JWToken := strings.Split(c.GetHeader("Authorization"), " ")[1]
	fmt.Println(JWToken)
	token, err := jwt.ParseWithClaims(
		JWToken,
		&util.CustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte("secretsecretsecretsecretsecretsecret"), nil
		},
	)
	if err != nil {
		fmt.Println(err)
		c.Abort()
		return
	}
	claims, ok := token.Claims.(*util.CustomClaims)
	if !ok {
		fmt.Println("Error")
	}
	if claims.ExpiresAt < time.Now().UTC().Unix() {
		fmt.Println("Token expired")
	}
	fmt.Println(claims.Username)

	c.Next()
}
