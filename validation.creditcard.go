package main

import (
	"net/http"
	"strconv"

	"github.com/go-validator/validator"
)

func validateCreditCard(w http.ResponseWriter, r *http.Request, creditcard CREDITCARD) bool {

	if errs := validator.Validate(creditcard); errs != nil {

		respondWith(w, r, nil, "Error(s) found in the credit card data: "+errs.Error(), nil, http.StatusBadRequest, false)
		return false

	}

	mm, _ := strconv.Atoi(creditcard.CardExpiryMM)

	if mm < 1 || mm > 12 {

		respondWith(w, r, nil, "Payment options's expiry month field should be between 1 and 12 (inclusive)", nil, http.StatusBadRequest, false)
		return false

	}

	return true

}
