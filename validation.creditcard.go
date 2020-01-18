package main

import (
	"net/http"
	"strconv"
)

func validateCreditCard(w http.ResponseWriter, r *http.Request, creditcard CREDITCARD) bool {

	if len(creditcard.Name) == 0 || len(creditcard.Name) > 100 {

		respondWith(w, r, nil, "Payment options's name field cannot be empty or greater than 100 characters long", nil, http.StatusBadRequest, false)
		return false

	}

	if len(creditcard.CardNumber) == 0 || len(creditcard.CardNumber) > 16 {

		respondWith(w, r, nil, "Payment options's card number field cannot be empty or greater than 16 characters long", nil, http.StatusBadRequest, false)
		return false

	}

	if len(creditcard.CardExpiryMM) == 0 || len(creditcard.CardExpiryMM) > 2 {

		respondWith(w, r, nil, "Payment options's expiry month field cannot be empty or greater than 2 characters long", nil, http.StatusBadRequest, false)
		return false

	}

	mm, _ := strconv.Atoi(creditcard.CardExpiryMM)

	if mm < 1 || mm > 12 {

		respondWith(w, r, nil, "Payment options's expiry month field should be between 1 and 12 (inclusive)", nil, http.StatusBadRequest, false)
		return false

	}

	if len(creditcard.CardExpiryYY) == 0 || len(creditcard.CardExpiryYY) > 2 {

		respondWith(w, r, nil, "Payment options's expiry year field cannot be empty or greater than 2 characters long", nil, http.StatusBadRequest, false)
		return false

	}

	if len(creditcard.SecurityCode) == 0 || len(creditcard.SecurityCode) > 3 {

		respondWith(w, r, nil, "Payment options's security code field cannot be empty or greater than 3 characters long", nil, http.StatusBadRequest, false)
		return false

	}

	if len(creditcard.ZipCode) == 0 || len(creditcard.ZipCode) > 10 {

		respondWith(w, r, nil, "Payment options's zip/pin code field cannot be empty or greater than 10 characters long", nil, http.StatusBadRequest, false)
		return false

	}

	return true

}
