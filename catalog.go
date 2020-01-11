package main

import (
	"encoding/json"
	"net/http"
)

func getProduct(w http.ResponseWriter, r *http.Request) {

	if !pre(w, r) {
		return
	}

}

func postProduct(w http.ResponseWriter, r *http.Request) {

	if !pre(w, r) {
		return
	}

	var p PRODUCT

	err := json.NewDecoder(r.Body).Decode(&p)

	if err != nil {
		respondWith(w, r, err, "Internal Error ...", nil, http.StatusBadRequest)
		return
	}

	insert("External", REDISCLIENT.Get(r.Header.Get("x-access-token")).Val()+".product", p)

	respondWith(w, r, nil, "Product Added ...", p, http.StatusCreated)

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
