package main

import (
	"net/http"
)

func validateCustomer(w http.ResponseWriter, r *http.Request, customer CUSTOMER) bool {

	if len(customer.FirstName) == 0 || len(customer.FirstName) > 100 {

		respondWith(w, r, nil, "Customer FirstName field cannot be empty or greater than 100 characters long", nil, http.StatusBadRequest, false)
		return false

	}

	if len(customer.LastName) == 0 || len(customer.LastName) > 100 {

		respondWith(w, r, nil, "Customer LastName field cannot be empty or greater than 100 characters long", nil, http.StatusBadRequest, false)
		return false

	}

	if !isValidEmail(customer.Email) {

		respondWith(w, r, nil, "Customer's email address need to be a valid email address", nil, http.StatusBadRequest, false)
		return false

	}

	if len(customer.Password) < 5 || len(customer.Password) > 1024 {

		respondWith(w, r, nil, "Customer's password cannot be less than 5 or more than 1024 characters", nil, http.StatusBadRequest, false)
		return false

	}

	if len(customer.PhoneNumbers) > 10 {

		respondWith(w, r, nil, "A customer object cannot have more than 100 phone numbers associated with it", nil, http.StatusBadRequest, false)
		return false

	}

	if len(customer.AddressBook) > 10 {

		respondWith(w, r, nil, "A customer object cannot have more than 100 addresses associated with it", nil, http.StatusBadRequest, false)
		return false

	}

	if len(customer.PaymentOptions) > 50 {

		respondWith(w, r, nil, "A customer object cannot have more than 50 PaymentOptions associated with it", nil, http.StatusBadRequest, false)
		return false

	}

	if len(customer.WishList) > 1000 {

		respondWith(w, r, nil, "A customer object cannot have more than 1000 wishlist skus associated with it", nil, http.StatusBadRequest, false)
		return false

	}

	if len(customer.SaveForLater) > 1000 {

		respondWith(w, r, nil, "A customer object cannot have more than 1000 SaveForLater skus associated with it", nil, http.StatusBadRequest, false)
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
