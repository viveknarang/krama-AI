package main

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

func authenticate(tokenString string) bool {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		secret := []byte("erjejkr48308dkfdjsfkldsj9048340958kjfklsdjf934403884309248ekjklfjflksjflkjklrjrjt485908539405kfjsdklfjsdklfjkljsfhghtrotu5turgmgf")

		return secret, nil

	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["foo"], claims["nbf"])
		return true
	} else {
		fmt.Println(err)
	}

	return false

}
