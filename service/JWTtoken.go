package service

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// var mySigningKey = os.Get("MY_JWT_TOKEN")
var mySigningKey = []byte("mysupersecretphrase")

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = "Matthew Bennett"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something went wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}
