package main

import (
	"encoding/json"
	"net/http"
	"net/url"
	"reflect"
)

func isValidJSON(s string) bool {

	var js interface{}
	return json.Unmarshal([]byte(s), &js) == nil

}

func isValidURL(toTest string) bool {

	_, err := url.ParseRequestURI(toTest)

	return !(err != nil)

}

func typeof(value interface{}) string {

	return reflect.TypeOf(value).String()

}

func pre(w http.ResponseWriter, r *http.Request) bool {

	if (r.Method == http.MethodPost || r.Method == http.MethodPut) && r.Header.Get("Content-Type") != "application/json" {

		respondWith(w, r, nil, MissingContentType, nil, http.StatusBadRequest)
		return false

	}

	if r.Header.Get("x-access-token") == "" {

		respondWith(w, r, nil, MissingAccessToken, nil, http.StatusBadRequest)
		return false

	}

	if !authenticate(r.Header.Get("x-access-token")) {

		respondWith(w, r, nil, InvalidSessionMessage, nil, http.StatusUnauthorized)
		return false

	}

	return true

}
