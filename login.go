package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func login(w http.ResponseWriter, r *http.Request) {

	var lc LOGIN

	err := json.NewDecoder(r.Body).Decode(&lc)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	results := find("Internal", "Customers", r.Body)

	for i := 0; i < len(results); i++ {
		fmt.Printf("%+v\n", *results[i])
	}

	currentTime := time.Now().Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"xkcd": "",
		"exp":  currentTime + 80000,
		"iat":  currentTime,
		"nbf":  currentTime - 100,
	})

	tokenString, err := token.SignedString([]byte("erjejkr48308dkfdjsfkldsj9048340958kjfklsdjf934403884309248ekjklfjflksjflkjklrjrjt485908539405kfjsdklfjsdklfjkljsfhghtrotu5turgmgf"))

	fmt.Println(tokenString, err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"token": ` + tokenString + ` }`))

}
