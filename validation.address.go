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

	if len(address.AddressLineOne) == 0 || len(address.AddressLineOne) > 200 {

		respondWith(w, r, nil, "AddressLineOne field cannot be empty or greater than 200 characters long", nil, http.StatusBadRequest, false)
		return false

	}

	if len(address.AddressLineTwo) > 200 {

		respondWith(w, r, nil, "AddressLineTwo field cannot be greater than 200 characters long", nil, http.StatusBadRequest, false)
		return false

	}

	if len(address.City) == 0 || len(address.City) > 100 {

		respondWith(w, r, nil, "Address city field cannot be empty or greater than 100 characters long", nil, http.StatusBadRequest, false)
		return false

	}

	if len(address.State) == 0 || len(address.State) > 100 {

		respondWith(w, r, nil, "Address state field cannot be empty or greater than 100 characters long", nil, http.StatusBadRequest, false)
		return false

	}

	if len(address.Country) == 0 || len(address.Country) > 100 {

		respondWith(w, r, nil, "Address country field cannot be empty or greater than 100 characters long", nil, http.StatusBadRequest, false)
		return false

	}

	if len(address.Pincode) == 0 || len(address.Pincode) > 10 {

		respondWith(w, r, nil, "Address pincode field cannot be empty or greater than 10 characters long", nil, http.StatusBadRequest, false)
		return false

	}

	return true

}
