package main

import (
	"net/http"

	"github.com/go-validator/validator"
)

func validateLoginRequest(w http.ResponseWriter, r *http.Request, loginRq LOGIN) bool {

	if errs := validator.Validate(loginRq); errs != nil {

		respondWith(w, r, nil, "Error(s) found in the login request: "+errs.Error(), nil, http.StatusBadRequest, false)
		return false

	}

	return true

}
