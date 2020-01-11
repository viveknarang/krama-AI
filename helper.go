package main

import (
	"encoding/json"
	"net/http"
)

func isValidJSON(s string) bool {

	var js interface{}
	return json.Unmarshal([]byte(s), &js) == nil

}

func pre(w http.ResponseWriter, r *http.Request) bool {

	if !authenticate(r.Header.Get("x-access-token")) {

		respondWith(w, r, nil, "Need to login first OR access token expired or invalid...", nil, http.StatusUnauthorized)
		return false

	}

	return true

}
