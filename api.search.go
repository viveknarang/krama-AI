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
		respondWith(w, r, err, HTTPBadRequestMessage, nil, http.StatusBadRequest, false)
		return
	}

	cidb := REDISCLIENT.Get(r.Header.Get("x-access-token")).Val()
	index := cidb + ProductGroupExtension + SearchIndexExtension

	searchResponse := quickSearch(index, sq.From, sq.To, sq.Query, sq.QueryFields, sq.ResponseFields)

	hits := make(map[int]interface{})
	results := make(map[string]interface{})

	for index, response := range searchResponse.Hits {
		hits[index] = response.Source
	}

	results["count"] = searchResponse.TotalHits.Value
	results["results"] = hits

	respondWith(w, r, nil, "Search Result ...", results, http.StatusOK, true)

}
