package main

import (
	"fmt"

	"github.com/go-redis/redis"

	"github.com/dgrijalva/jwt-go"
)

func authenticate(tokenString string) bool {

	var isValid bool

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		secret := []byte(JWTSecret)

		return secret, nil

	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		if REDISCLIENT.Get(tokenString).Err() == redis.Nil {
			REDISCLIENT.Set(tokenString, claims["cxs"], 0)
		}

		isValid = true

	} else {

		fmt.Println(err)
		isValid = false

	}

	return isValid

}
