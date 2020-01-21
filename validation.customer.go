package main

import (
	"errors"
	"net/http"
	"reflect"
	"strconv"

	"github.com/go-validator/validator"
)

func customValidatorForEmail(v interface{}, param string) error {

	st := reflect.ValueOf(v)

	if !isValidEmail(st.String()) {
		return errors.New("field is not a valid email address")
	}

	return nil
}

func customValidatorForTypeArrayLengths(v interface{}, param string) error {

	st := reflect.ValueOf(v)

	iparam, _ := strconv.Atoi(param)

	if st.Len() > int(iparam) {
		return errors.New("field has more than " + param + " entries")
	}

	return nil

}

func validateCustomer(w http.ResponseWriter, r *http.Request, customer CUSTOMER) bool {

	validator.SetValidationFunc("validateEmail", customValidatorForEmail)
	validator.SetValidationFunc("validateTypeArrayLength", customValidatorForTypeArrayLengths)

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
