package main

import (
	"net/http"
	"strings"

	"github.com/romana/rlog"
)

func pre(w http.ResponseWriter, r *http.Request) bool {

	if (r.Method == http.MethodPost || r.Method == http.MethodPut) && r.Header.Get("Content-Type") != "application/json" {

		rlog.Debug("pre(): Missing content type ...")
		respondWith(w, r, nil, MissingContentType, nil, http.StatusBadRequest, false)
		return false

	}

	authToken := r.Header.Get("Authorization")

	if len(authToken) == 0 {

		rlog.Debug("pre(): Missing authorization in the header ...")
		respondWith(w, r, nil, MissingAuthRequestHeader, nil, http.StatusBadRequest, false)
		return false

	}

	if !strings.Contains(authToken, "Bearer") {

		rlog.Debug("pre(): Malformed authorization in the header ...")
		respondWith(w, r, nil, MalformedAuthRequestHeader, nil, http.StatusBadRequest, false)
		return false

	}

	splitToken := strings.Split(authToken, "Bearer ")
	authToken = splitToken[1]

	if authToken == "" {

		rlog.Debug("pre(): Missing access token ...")
		respondWith(w, r, nil, MissingAccessToken, nil, http.StatusBadRequest, false)
		return false

	}

	if !areCoreServicesUp() {

		rlog.Debug("pre(): Core services seems to be down ...")
		respondWith(w, r, nil, ServiceDownMessage, nil, http.StatusServiceUnavailable, false)
		return false

	}

	if !authenticate(authToken) {

		rlog.Debug("pre(): Authentication failed ...")
		respondWith(w, r, nil, InvalidSessionMessage, nil, http.StatusUnauthorized, false)
		return false

	}

	return true

}
