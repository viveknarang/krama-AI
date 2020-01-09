package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func login(w http.ResponseWriter, r *http.Request) {

	currentTime := time.Now().Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": currentTime + 120,
		"iat": currentTime,
		"nbf": currentTime - 100,
	})

	tokenString, err := token.SignedString([]byte("erjejkr48308dkfdjsfkldsj9048340958kjfklsdjf934403884309248ekjklfjflksjflkjklrjrjt485908539405kfjsdklfjsdklfjkljsfhghtrotu5turgmgf"))

	fmt.Println(tokenString, err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"token": ` + tokenString + ` }`))

}
