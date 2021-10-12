package controllers

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"restapiGin/environment"
	"restapiGin/service"
)

func Refresh(c *gin.Context) {
	mapToken := map[string]string{}
	if err := c.ShouldBindJSON(&mapToken); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	refreshToken := mapToken["refresh_token"]

	//verify the token
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		// Make sure the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(environment.ViperEnvVariable("JWT_REFRESH_SECRET")), nil
	})

	//if there is an error, the token must have expired
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Refresh token expired")
		return
	}

	//is token valid?
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		c.JSON(http.StatusUnauthorized, err)
		return
	}
	//Since token is valid, get the uuid:
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		refreshUuid, ok := claims["refresh_uuid"].(string)
		if !ok {
			c.JSON(http.StatusUnprocessableEntity, err)
			return
		}
		userId, err := claims["userid"].(string)
		if !err {
			c.JSON(http.StatusUnprocessableEntity, "Error occured")
			return
		}
		//Delete the previous Refresh Token
		deleted, delErr := service.DeleteAuth(refreshUuid)
		if delErr != nil || deleted == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			return
		}

		userUuid, _ := uuid.FromString(userId)

		//Create new pairs of refresh and access tokens
		ts, createErr := service.CreateToken(userUuid)
		if createErr != nil {
			c.JSON(http.StatusForbidden, createErr.Error())
			return
		}

		//save the tokens metadata to redis
		saveErr := service.CreateAuth(userUuid, ts)
		if saveErr != nil {
			c.JSON(http.StatusForbidden, saveErr.Error())
			return
		}
		tokens := map[string]string{
			"access_token":  ts.AccessToken,
			"refresh_token": ts.RefreshToken,
		}
		c.JSON(http.StatusCreated, tokens)
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "refresh expired"})
	}
}
