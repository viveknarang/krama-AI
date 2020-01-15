package main

import (
	"net/http"
)

func validateCreditCard(w http.ResponseWriter, r *http.Request, creditcard CREDITCARD) bool {

	if len(creditcard.Name) == 0 || len(creditcard.Name) > 100 {

		respondWith(w, r, nil, "Credit Card's Name field cannot be empty or greater than 100 characters long", nil, http.StatusBadRequest, false)
		return false

	}

	//TODO: More complex credit card field validations needed here!

	return true

}
