package main

import (
	"net/http"
)

func validateAddress(w http.ResponseWriter, r *http.Request, address ADDRESS) bool {

	if len(address.FirstName) == 0 || len(address.FirstName) > 100 {

		respondWith(w, r, nil, "Address FirstName field cannot be empty or greater than 100 characters long", nil, http.StatusBadRequest, false)
		return false

	}

	if len(address.LastName) == 0 || len(address.LastName) > 100 {

		respondWith(w, r, nil, "Address LastName field cannot be empty or greater than 100 characters long", nil, http.StatusBadRequest, false)
		return false

	}

	// TODO: More complex validations needed here!

	return true

}
