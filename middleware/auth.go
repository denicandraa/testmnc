package middleware

import (
	"../helpers"
	"../models"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

type MiddleWare struct {
	h helpers.HelpersJson
}

var jwtKey = []byte("mncbank")

func (u *MiddleWare)CreateAuth(clientUsername string, clientRole string) (token string, err error) {
	sign := jwt.New(jwt.SigningMethodHS256)
	claims := sign.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["client_username"] = clientUsername
	claims["client_role"] = clientRole
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenx, errx := sign.SignedString(jwtKey)

	data := models.Auth{User_name: clientUsername,Token: tokenx}
	u.h.WriteFileJsonAuthLogin(data)
	return tokenx, errx
}

func (u *MiddleWare)Auth(c *gin.Context) {

	tokenString := c.Request.Header.Get("Authorization")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return jwtKey, nil
	})

	if token != nil && err == nil {

		var akses bool
		dataAuth := u.h.GetAuth()

		for i := range dataAuth {
			if dataAuth[i].Token == tokenString{
				akses = true
				break
			}
		}

		if !akses{
			result := models.Response{}

			result.ApiMessage = "Sesi anda telah habis , silahkan login ulang"
			result.Data = nil

			c.JSON(403, result)
			c.Abort()
		}

		c.Next()

	} else {

		result := models.Response{}

		result.ApiMessage = "Sesi anda telah habis , silahkan login ulang"
		result.Data = nil

		c.JSON(403, result)
		c.Abort()
	}

}
