package main

import "net/http"

func pre(w http.ResponseWriter, r *http.Request) bool {

	if (r.Method == http.MethodPost || r.Method == http.MethodPut) && r.Header.Get("Content-Type") != "application/json" {

		respondWith(w, r, nil, MissingContentType, nil, http.StatusBadRequest)
		return false

	}

	if r.Header.Get("x-access-token") == "" {

		respondWith(w, r, nil, MissingAccessToken, nil, http.StatusBadRequest)
		return false

	}

	if !areCoreServicesUp() {

		respondWith(w, r, nil, ServiceDownMessage, nil, http.StatusServiceUnavailable)
		return false

	}

	if !authenticate(r.Header.Get("x-access-token")) {

		respondWith(w, r, nil, InvalidSessionMessage, nil, http.StatusUnauthorized)
		return false

	}

	return true

}
