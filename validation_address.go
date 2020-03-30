package main

import (
	"net/http"

	"github.com/go-validator/validator"
)

func validateAddress(w http.ResponseWriter, r *http.Request, address ADDRESS) bool {

	if errs := validator.Validate(address); errs != nil {

		respondWith(w, r, nil, "Error(s) found in the address data: "+errs.Error(), nil, http.StatusBadRequest, false)
		return false

	}

	return true

}
