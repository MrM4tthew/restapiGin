package service

import (
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
	"restapiGin/environment"
	"restapiGin/models"
	"time"
)

//var mySigningKey = []byte(environment.ViperEnvVariable("JWT_SECRET"))
//
//func CreateToken(username string, email string) (string, error) {
//	token := jwt.New(jwt.SigningMethodHS256)
//
//	claims := token.Claims.(jwt.MapClaims)
//
//	claims["username"] = username
//	claims["email"] = email
//	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
//
//	tokenString, err := token.SignedString(mySigningKey)
//
//	if err != nil {
//		fmt.Errorf("Something went wrong: %s", err.Error())
//		return "", err
//	}
//
//	return tokenString, nil
//}

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}

func CreateToken(userid uuid.UUID) (*TokenDetails, error) {
	td := &TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	td.AccessUuid = uuid.NewV4().String()

	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUuid = uuid.NewV4().String()

	var err error

	// Creating access token
	atClaims := jwt.MapClaims{}
	atClaims["access_uuid"] = td.AccessUuid
	atClaims["userid"] = userid
	atClaims["exp"] = td.AtExpires
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = token.SignedString([]byte(environment.ViperEnvVariable("JWT_SECRET")))
	if err != nil {
		return nil, err
	}

	// Creating refresh token
	rtClaims := token.Claims.(jwt.MapClaims)
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["userid"] = userid
	rtClaims["exp"] = td.RtExpires
	td.RefreshToken, err = token.SignedString([]byte(environment.ViperEnvVariable("JWT_REFRESH_SECRET")))
	if err != nil {
		return nil, err
	}

	return td, nil
}

func CreateAuth(userid uuid.UUID, td *TokenDetails) error {
	at := time.Unix(td.AtExpires, 0)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	errAccess := models.Client.Set(td.AccessUuid, userid.String(), at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}
	errRefresh := models.Client.Set(td.RefreshUuid, userid.String(), rt.Sub(now)).Err()
	if errRefresh != nil {
		return errRefresh
	}

	return nil
}
