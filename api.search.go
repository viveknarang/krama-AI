package main

import "net/http"

func basicProductGroupSearch(w http.ResponseWriter, r *http.Request) {

	if !pre(w, r) {
		return
	}

}
