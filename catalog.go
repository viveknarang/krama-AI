package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func pre(w http.ResponseWriter, r *http.Request) bool {

	if !authenticate(r.Header.Get("x-access-token")) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"message": "Need to login first OR access token expired or invalid..."}`))
		return false
	}

	return true

}

func getProduct(w http.ResponseWriter, r *http.Request) {

	if !pre(w, r) {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "get called"}`))
}

func postProduct(w http.ResponseWriter, r *http.Request) {

	if !pre(w, r) {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	var p PRODUCT

	err := json.NewDecoder(r.Body).Decode(&p)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	insert("test", "test", p)

	var response RESPONSE

	response.Code = "200"
	response.Message = "Product Added ..."
	response.Success = true
	response.Time = time.Now().Unix()
	response.Response = p

	resp, err := json.Marshal(response)

	w.Write([]byte(resp))

}

func putProduct(w http.ResponseWriter, r *http.Request) {

	if !pre(w, r) {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(`{"message": "put called"}`))
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {

	if !pre(w, r) {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "delete called"}`))
}

func notFound(w http.ResponseWriter, r *http.Request) {

	if !pre(w, r) {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "not found"}`))
}
