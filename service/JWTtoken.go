package service

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"restapiGin/environment"
	"time"
)

var mySigningKey = environment.ViperEnvVariable("JWT_SECRET")

func CreateToken(username string, email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["username"] = username
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something went wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

//type TokenDetails struct {
//	AccessToken  string
//	RefreshToken string
//	AccessUuid   string
//	RefreshUuid  string
//	AtExpires    int64
//	RtExpires    int64
//}
//
//func CreateToken(userid uuid.UUID, username string, email string) (*TokenDetails, error) {
//	td := &TokenDetails{}
//	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
//	td.AccessUuid = uuid.NewV4().String()
//
//	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
//	td.RefreshUuid = uuid.NewV4().String()
//
//	var err error
//	token := jwt.New(jwt.SigningMethodHS256)
//	// Creating access token
//	atClaims := token.Claims.(jwt.MapClaims)
//	atClaims["access_uuid"] = td.AccessUuid
//	atClaims["userid"] = userid
//	atClaims["username"] = username
//	atClaims["email"] = email
//	atClaims["exp"] = td.AtExpires
//	td.AccessToken, err = token.SignedString(environment.ViperEnvVariable("JWT_SECRET"))
//	if err != nil {
//		return nil, err
//	}
//
//	// Creating refresh token
//	rtClaims := token.Claims.(jwt.MapClaims)
//	rtClaims["refresh_uuid"] = td.RefreshUuid
//	rtClaims["userid"] = userid
//	rtClaims["username"] = username
//	rtClaims["email"] = email
//	rtClaims["exp"] = td.RtExpires
//	td.RefreshToken, err = token.SignedString(environment.ViperEnvVariable("JWT_REFRESH_SECRET"))
//	if err != nil {
//		return nil, err
//	}
//
//	return td, nil
//}
