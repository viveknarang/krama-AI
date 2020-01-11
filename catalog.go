package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/bson"
)

func getProduct(w http.ResponseWriter, r *http.Request) {

	if !pre(w, r) {
		return
	}

	var jx []byte

	redisC := REDISCLIENT.Get(r.URL.Path)

	if redisC.Err() != redis.Nil {

		jx = []byte(redisC.Val())

	} else {

		pth := strings.Split(r.URL.Path, "/")
		sku := pth[len(pth)-1]

		dbcol := REDISCLIENT.Get(r.Header.Get("x-access-token")).Val() + ProductExtension

		results := find(ExternalDB, dbcol, bson.M{"sku": sku})

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

		REDISCLIENT.Set(r.URL.Path, j, 0)

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

	p.Updated = time.Now().UnixNano()

	dbcol := REDISCLIENT.Get(r.Header.Get("x-access-token")).Val() + ProductExtension

	insert(ExternalDB, dbcol, p)

	if syncProductGroup(w, r, p) {

		respondWith(w, r, nil, ProductAddedMessage, p, http.StatusCreated)

	} else {

		respondWith(w, r, nil, ProductNotAddedMessage, p, http.StatusNotModified)

	}

}

func putProduct(w http.ResponseWriter, r *http.Request) {

	if !pre(w, r) {
		return
	}

	var p PRODUCT

	err := json.NewDecoder(r.Body).Decode(&p)

	if err != nil {
		respondWith(w, r, err, HTTPBadRequestMessage, nil, http.StatusBadRequest)
		return
	}

	dbcol := REDISCLIENT.Get(r.Header.Get("x-access-token")).Val() + ProductExtension

	result := update(ExternalDB, dbcol, bson.M{"sku": p.Sku}, bson.M{"$set": p})

	if result[0] == 1 && result[1] == 1 {

		if syncProductGroup(w, r, p) {

			REDISCLIENT.Del(r.URL.Path)
			respondWith(w, r, nil, ProductUpdatedMessage, p, http.StatusAccepted)

		} else {

			respondWith(w, r, nil, ProductNotUpdatedMessage, p, http.StatusNotModified)

		}

	} else if result[0] == 1 && result[1] == 0 {

		respondWith(w, r, nil, ProductNotUpdatedMessage, nil, http.StatusNotModified)

	} else if result[0] == 0 && result[1] == 0 {

		respondWith(w, r, nil, ProductNotFoundMessage, nil, http.StatusNotModified)

	}

}

func deleteProduct(w http.ResponseWriter, r *http.Request) {

	if !pre(w, r) {
		return
	}

	dbcol := REDISCLIENT.Get(r.Header.Get("x-access-token")).Val() + ProductExtension

	pth := strings.Split(r.URL.Path, "/")
	sku := pth[len(pth)-1]

	results := find(ExternalDB, dbcol, bson.M{"sku": sku})

	if len(results) != 1 {
		respondWith(w, r, nil, ProductNotFoundMessage, nil, http.StatusNotFound)
		return
	}

	j, err0 := bson.MarshalExtJSON(results[0], false, false)

	if err0 != nil {
		respondWith(w, r, err0, HTTPInternalServerErrorMessage, nil, http.StatusInternalServerError)
		return
	}

	var product PRODUCT

	err1 := json.Unmarshal([]byte(j), &product)

	if err1 != nil {
		respondWith(w, r, err1, HTTPInternalServerErrorMessage, nil, http.StatusInternalServerError)
		return
	}

	if deleteDocument(ExternalDB, dbcol, bson.M{"sku": sku}) == 1 {

		if syncProductGroup(w, r, product) {

			REDISCLIENT.Del(r.URL.Path)
			respondWith(w, r, nil, ProductDeletedMessage, nil, http.StatusOK)

		} else {

			respondWith(w, r, nil, ProductNotDeletedMessage, nil, http.StatusOK)

		}

	} else {

		respondWith(w, r, nil, ProductNotFoundMessage, nil, http.StatusNotModified)

	}

}

func getProductGroup(w http.ResponseWriter, r *http.Request) {

	if !pre(w, r) {
		return
	}

	var jx []byte

	redisC := REDISCLIENT.Get(r.URL.Path)

	if redisC.Err() != redis.Nil {

		jx = []byte(redisC.Val())

	} else {

		dbcol := REDISCLIENT.Get(r.Header.Get("x-access-token")).Val() + ProductGroupExtension

		pth := strings.Split(r.URL.Path, "/")
		pgid := pth[len(pth)-1]

		results := find(ExternalDB, dbcol, bson.M{"groupid": pgid})

		if len(results) != 1 {
			respondWith(w, r, nil, ProductGroupNotFoundMessage, nil, http.StatusNotFound)
			return
		}

		j, err0 := bson.MarshalExtJSON(results[0], false, false)

		if err0 != nil {
			respondWith(w, r, err0, HTTPInternalServerErrorMessage, nil, http.StatusInternalServerError)
			return
		}

		jx = j

		REDISCLIENT.Set(r.URL.Path, j, 0)

	}

	var productGroup PRODUCTGROUP

	err1 := json.Unmarshal([]byte(jx), &productGroup)

	if err1 != nil {
		respondWith(w, r, err1, HTTPInternalServerErrorMessage, nil, http.StatusInternalServerError)
		return
	}

	respondWith(w, r, nil, ProductGroupFoundMessage, productGroup, http.StatusOK)

}

func deleteProductGroup(w http.ResponseWriter, r *http.Request) {

	if !pre(w, r) {
		return
	}

	dbcol := REDISCLIENT.Get(r.Header.Get("x-access-token")).Val() + ProductGroupExtension

	pth := strings.Split(r.URL.Path, "/")
	pgid := pth[len(pth)-1]

	if deleteDocument(ExternalDB, dbcol, bson.M{"groupid": pgid}) == 1 {
		REDISCLIENT.Del(r.URL.Path)
		respondWith(w, r, nil, ProductGroupDeletedMessage, nil, http.StatusOK)
	} else {
		respondWith(w, r, nil, ProductGroupNotFoundMessage, nil, http.StatusNotModified)
	}

}
