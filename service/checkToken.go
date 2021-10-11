package service

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"restapiGin/environment"
	"restapiGin/models"
	"strings"
)

type AccesDetails struct {
	AccessUuid string
	UserId     string
}

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(environment.ViperEnvVariable("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func TokenValid(r *http.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

func ExtractTokenMetadata(r *http.Request) (*AccesDetails, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok2 := claims["access_uuid"].(string)
		if !ok2 {
			return nil, err
		}
		userId, ok3 := claims["userid"].(string)
		if !ok3 {
			return nil, err
		}

		return &AccesDetails{
			AccessUuid: accessUuid,
			UserId:     userId,
		}, nil
	}

	return nil, err
}

func FetchAuth(authD *AccesDetails) (string, error) {
	userid, err := models.Client.Get(authD.AccessUuid).Result()
	if err != nil {
		return "", err
	}

	//userId, _ := strconv.ParseUint(userid, 10, 64)

	return userid, nil
}

func DeleteAuth(givenUuid string) (int64, error) {
	deleted, err := models.Client.Del(givenUuid).Result()
	if err != nil {
		return 0, err
	}

	return deleted, nil
}
