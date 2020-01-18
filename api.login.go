package main

import (
	"encoding/json"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"github.com/romana/rlog"
)

func login(w http.ResponseWriter, r *http.Request) {

	rlog.Debug("login() handle function invoked ...")

	var rx LOGIN

	err := json.NewDecoder(r.Body).Decode(&rx)

	if err != nil {

		respondWith(w, r, err, HTTPBadRequestMessage, nil, http.StatusBadRequest, false)
		return

	}

	if !areCoreServicesUp() {

		respondWith(w, r, nil, ServiceDownMessage, nil, http.StatusServiceUnavailable, false)
		return

	}

	results := findMongoDocument(InternalDB, CustomersDB, bson.M{"CustomerID": rx.CustomerID, "APIKey": rx.APIKey})

	if len(results) != 1 {

		respondWith(w, r, nil, LoginFailedMessage, nil, http.StatusUnauthorized, false)

	} else {

		var customer PLATFORMCUSTOMER

		j, err0 := bson.MarshalExtJSON(results[0], false, false)

		if err0 != nil {
			respondWith(w, r, err0, HTTPInternalServerErrorMessage, nil, http.StatusInternalServerError, false)
			return
		}

		err1 := json.Unmarshal([]byte(j), &customer)

		if err1 != nil {
			respondWith(w, r, err1, HTTPInternalServerErrorMessage, nil, http.StatusInternalServerError, false)
			return
		}

		currentTime := time.Now().Unix()

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"cxs": customer.Secret,
			"uid": uuid.New(),
			"exp": currentTime + LoginSessionDuration,
			"iat": currentTime,
			"nbf": currentTime - 100,
		})

		tokenString, err := token.SignedString([]byte(JWTSecret))

		if REDISCLIENT.Get(tokenString).Err() == redis.Nil {
			REDISCLIENT.Set(tokenString, customer.Secret, 0)
		}

		respondWith(w, r, err, LoginSuccessMessage, bson.M{"Token": tokenString, "ValidForSeconds": LoginSessionDuration}, http.StatusOK, true)

	}

}
