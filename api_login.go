package main

import (
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"github.com/romana/rlog"
)

func login(w http.ResponseWriter, r *http.Request) {

	rlog.Debug("login() handle function invoked ...")

	var rx LOGIN

	if !mapInput(w, r, &rx) {
		return
	}

	if !validateLoginRequest(w, r, rx) {
		return
	}

	if !areCoreServicesUp() {

		respondWith(w, r, nil, ServiceDownMessage, nil, http.StatusServiceUnavailable, false)
		return

	}

	var opts options.FindOptions

	results := findMongoDocument(InternalDB, CustomersDB, bson.M{"CustomerID": rx.CustomerID, "APIKey": rx.APIKey}, &opts)

	if len(results) != 1 {

		respondWith(w, r, nil, LoginFailedMessage, nil, http.StatusUnauthorized, false)

	} else {

		var customer PLATFORMCUSTOMER

		mapDocument(w, r, &customer, results[0])

		currentTime := time.Now().Unix()

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"cxs": customer.Secret,
			"rxt": customer.Rate,
			"uid": uuid.New(),
			"exp": currentTime + LoginSessionDuration,
			"iat": currentTime,
			"nbf": currentTime - 100,
		})

		tokenString, err := token.SignedString([]byte(JWTSecret))

		if REDISCLIENT.Get(tokenString).Err() == redis.Nil {
			REDISCLIENT.Set(tokenString, customer.Secret, 0)
		}

		if REDISCLIENT.Get(tokenString+"_rxt").Err() == redis.Nil {
			REDISCLIENT.Set(tokenString+"_rxt", customer.Rate, 0)
		}

		respondWith(w, r, err, LoginSuccessMessage, bson.M{"Token": tokenString, "ValidForSeconds": LoginSessionDuration}, http.StatusOK, true)

	}
	

}
