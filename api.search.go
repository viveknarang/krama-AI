package main

import (
	"encoding/json"
	"net/http"
)

func basicProductGroupSearch(w http.ResponseWriter, r *http.Request) {

	if !pre(w, r) {
		return
	}

	var sq SEARCHREQUEST

	err := json.NewDecoder(r.Body).Decode(&sq)

	if err != nil {
		respondWith(w, r, err, HTTPBadRequestMessage, nil, http.StatusBadRequest)
		return
	}

	cidb := REDISCLIENT.Get(r.Header.Get("x-access-token")).Val()
	index := cidb + ProductGroupExtension + SearchIndexExtension

	searchResponse := queryES(index, 0, 100, sq.Q, sq.Fields)

	respondWith(w, r, nil, "Search Result ...", searchResponse, http.StatusOK)

}
