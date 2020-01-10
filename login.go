package main

import (
	"encoding/json"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/dgrijalva/jwt-go"
)

func login(w http.ResponseWriter, r *http.Request) {

	var rx LOGIN

	err := json.NewDecoder(r.Body).Decode(&rx)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	results := find("Internal", "Customers", bson.M{"CustomerID": rx.CustomerID, "APIKey": rx.APIKey})

	if len(results) != 1 {

		respondWith(w, r, nil, "Invalid Login ...", nil, http.StatusUnauthorized)

	} else {

		currentTime := time.Now().Unix()

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"xkcd": "",
			"exp":  currentTime + 80000,
			"iat":  currentTime,
			"nbf":  currentTime - 100,
		})

		tokenString, err := token.SignedString([]byte("erjejkr48308dkfdjsfkldsj9048340958kjfklsdjf934403884309248ekjklfjflksjflkjklrjrjt485908539405kfjsdklfjsdklfjkljsfhghtrotu5turgmgf"))

		respondWith(w, r, err, "Login Successful...", bson.M{"token": tokenString}, http.StatusOK)

	}

}
