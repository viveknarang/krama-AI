package main

import "net/http"

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
	w.Write([]byte(`{"message": "post called"}`))
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
