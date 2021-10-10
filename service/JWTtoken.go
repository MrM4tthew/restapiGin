package service

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"restapiGin/environment"
	"time"
)

var mySigningKey = environment.ViperEnvVariable("JWT_SECRET")

func GenerateJWT(username string, email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["username"] = username
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something went wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}
