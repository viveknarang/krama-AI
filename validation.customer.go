package main

import (
	"net/http"
)

func validateCustomer(w http.ResponseWriter, r *http.Request, customer CUSTOMER) bool {

	if len(customer.FirstName) == 0 || len(customer.FirstName) > 100 {

		respondWith(w, r, nil, "Customer FirstName field cannot be empty or greater than 100 characters long", nil, http.StatusBadRequest, false)
		return false

	}

	return true

}
