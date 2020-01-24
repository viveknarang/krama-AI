package main

import (
	"net/http"

	"github.com/go-validator/validator"
)

func validateCustomer(w http.ResponseWriter, r *http.Request, customer CUSTOMER) bool {

	if errs := validator.Validate(customer); errs != nil {

		respondWith(w, r, nil, "Error(s) found in the customer data: "+errs.Error(), nil, http.StatusBadRequest, false)
		return false

	}

	for _, address := range customer.AddressBook {
		if !validateAddress(w, r, address) {
			return false
		}
	}

	for _, paymentOption := range customer.PaymentOptions {
		if !validateCreditCard(w, r, paymentOption) {
			return false
		}
	}

	return true

}
