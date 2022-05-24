package main

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte(CONFIG["TOKEN_KEY"])

func generateJwtToken(userModel User) (string, error) {
	expiresAt := time.Now().Local().Add(24 * time.Hour).Unix()
	jwtClaims := jwt.MapClaims{}
	jwtClaims["user"] = map[string]interface{}{
		"id":       userModel.ID,
		"username": userModel.Username,
		"email":    userModel.Email,
		"role":     userModel.Role,
	}
	jwtClaims["expiresAt"] = expiresAt

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)

	jwtTokenString, err := jwtToken.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return jwtTokenString, nil
}

func validateJwtToken(jwtTokenString string) (interface{}, error) {
	jwtToken, err := jwt.Parse(
		jwtTokenString,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("failed to validate and sign the token")
			}

			return jwtKey, nil
		},
	)

	if err != nil {
		return nil, err
	}

	if !jwtToken.Valid {
		return nil, errors.New("invalid token")
	}

	return jwtToken, nil
}

func getJwtToken(c *gin.Context) (string, bool) {
	requestHeader := c.GetHeader("Authorization")
	requestHeaderArr := strings.Split(requestHeader, " ")
	if len(requestHeaderArr) != 2 {
		return "", false
	}

	authRequestType := strings.Trim(requestHeaderArr[0], "\n\r\t")
	if strings.ToLower(authRequestType) != strings.ToLower("Bearer") {
		return "", false
	}

	return strings.Trim(requestHeaderArr[1], "\n\t\r"), true
}
