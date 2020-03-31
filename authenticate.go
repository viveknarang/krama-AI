package main

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/romana/rlog"

	"github.com/dgrijalva/jwt-go"
)

func authenticate(tokenString string) bool {

	rlog.Debug("authenticate() function invoked for token: " + tokenString)

	var isValid bool

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		secret := []byte(JWTSecret)

		return secret, nil

	})

	if err != nil {
		return false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		if REDISCLIENT.Get(tokenString).Err() == redis.Nil {
			REDISCLIENT.Set(tokenString, claims["cxs"], 0)
		}

		if REDISCLIENT.Get(tokenString+"_rxt").Err() == redis.Nil {
			REDISCLIENT.Set(tokenString+"_rxt", claims["rxt"], 0)
		}

		isValid = true

	} else {

		if REDISCLIENT.Get(tokenString).Err() != redis.Nil {
			REDISCLIENT.Del(tokenString)
		}

		if REDISCLIENT.Get(tokenString+"_rxt").Err() != redis.Nil {
			REDISCLIENT.Del(tokenString + "_rxt")
		}

		isValid = false

	}

	return isValid

}
