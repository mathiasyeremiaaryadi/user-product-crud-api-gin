package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte(CONFIG["TOKEN_KEY"])

func generateJwtToken(userModel User) (string, error) {
	expiresAt := time.Now().Local().Add(24 * time.Hour).Unix()
	jwtClaims := JWTClaims{
		StandardClaims: jwt.StandardClaims{
			Subject:   userModel.Username,
			ExpiresAt: expiresAt,
		},
		user: userModel,
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)

	jwtTokenString, err := jwtToken.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return jwtTokenString, nil
}

func validateJwtToken(jwtTokenString string) (interface{}, error) {
	var jwtClaims JWTClaims
	jwtToken, err := jwt.ParseWithClaims(
		jwtTokenString,
		&jwtClaims,
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

	return jwtClaims.user, nil
}
