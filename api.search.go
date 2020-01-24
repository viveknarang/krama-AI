package main

import (
	"net/http"

	"github.com/romana/rlog"
)

func quickSearch(w http.ResponseWriter, r *http.Request) {

	rlog.Debug("quickSearch() handle function invoked ...")

	if !pre(w, r) {
		return
	}

	var sq SEARCHREQUEST

	if !mapInput(w, r, &sq) {
		return
	}

	cidb := getAccessToken(r)
	index := cidb + ProductGroupExtension + SearchIndexExtension

	searchResponse := basicSearch(index, sq.From, sq.To, sq.Query, sq.QueryFields, sq.ResponseFields)

	hits := make(map[int]interface{})
	results := make(map[string]interface{})

	for index, response := range searchResponse.Hits {
		hits[index] = response.Source
	}

	results["count"] = searchResponse.TotalHits.Value
	results["results"] = hits

	respondWith(w, r, nil, "Search Result ...", results, http.StatusOK, true)

}

func fullpageSearch(w http.ResponseWriter, r *http.Request) {

	rlog.Debug("fullpageSearch() handle function invoked ...")

	if !pre(w, r) {
		return
	}

	var sq SEARCHREQUEST

	if !mapInput(w, r, &sq) {
		return
	}

	cidb := getAccessToken(r)
	index := cidb + ProductGroupExtension + SearchIndexExtension

	searchResponse := facetedSearch(index, sq.From, sq.To, sq.Query, sq.QueryFields, sq.ResponseFields, sq.TermFacetFields, sq.RangeFacetFields)

	results := make(map[string]interface{})

	results["facets"] = searchResponse.Aggregations
	results["hits"] = searchResponse.Hits.Hits
	results["count"] = searchResponse.Hits.TotalHits.Value

	respondWith(w, r, nil, "Search Result ...", results, http.StatusOK, true)

}
