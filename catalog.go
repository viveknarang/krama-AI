package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/bson"
)

func getProduct(w http.ResponseWriter, r *http.Request) {

	if !pre(w, r) {
		return
	}

	var jx []byte

	redisC := REDISCLIENT.Get(r.Method + ":" + r.URL.Path)

	if redisC.Err() != redis.Nil {

		jx = []byte(redisC.Val())

	} else {

		pth := strings.Split(r.URL.Path, "/")
		sku := pth[len(pth)-1]

		results := find(ExternalDB, REDISCLIENT.Get(r.Header.Get("x-access-token")).Val()+ProductExtension, bson.M{"sku": sku})

		if len(results) != 1 {
			respondWith(w, r, nil, ProductNotFoundMessage, nil, http.StatusNotFound)
			return
		}

		j, err0 := bson.MarshalExtJSON(results[0], false, false)

		if err0 != nil {
			respondWith(w, r, err0, HTTPInternalServerErrorMessage, nil, http.StatusInternalServerError)
			return
		}

		jx = j

		REDISCLIENT.Set(r.Method+":"+r.URL.Path, j, 0)

	}

	var product PRODUCT

	err1 := json.Unmarshal([]byte(jx), &product)

	if err1 != nil {
		respondWith(w, r, err1, HTTPInternalServerErrorMessage, nil, http.StatusInternalServerError)
		return
	}

	respondWith(w, r, nil, ProductFoundMessage, product, http.StatusOK)

}

func postProduct(w http.ResponseWriter, r *http.Request) {

	if !pre(w, r) {
		return
	}

	var p PRODUCT

	err := json.NewDecoder(r.Body).Decode(&p)

	if err != nil {
		respondWith(w, r, err, HTTPBadRequestMessage, nil, http.StatusBadRequest)
		return
	}

	insert(ExternalDB, REDISCLIENT.Get(r.Header.Get("x-access-token")).Val()+ProductExtension, p)

	respondWith(w, r, nil, ProductAddedMessage, p, http.StatusCreated)

}

func putProduct(w http.ResponseWriter, r *http.Request) {

	if !pre(w, r) {
		return
	}

}

func deleteProduct(w http.ResponseWriter, r *http.Request) {

	if !pre(w, r) {
		return
	}

}

func notFound(w http.ResponseWriter, r *http.Request) {

	if !pre(w, r) {
		return
	}

}
