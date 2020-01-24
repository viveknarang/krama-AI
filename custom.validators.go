package main

import (
	"errors"
	"math"
	"reflect"
	"strconv"
	"strings"

	"github.com/go-validator/validator"
)

func setCustomValidators() {

	validator.SetValidationFunc("size", customValidatorForSize)
	validator.SetValidationFunc("hasNoSpaces", customValidatorForNoSpaces)
	validator.SetValidationFunc("checkMaxFloat", customValidatorForMaxFloat)
	validator.SetValidationFunc("isValidCurrency", customValidatorForAllowedCurrencies)
	validator.SetValidationFunc("validateEmail", customValidatorForEmail)
	validator.SetValidationFunc("validateTypeArrayLength", customValidatorForSize)
	validator.SetValidationFunc("validStarRating", customValidatorForStarRating)
	validator.SetValidationFunc("isValidSortOrder", customValidatorForSortOrder)

}

func customValidatorForSize(v interface{}, param string) error {

	st := reflect.ValueOf(v)

	iparam, _ := strconv.Atoi(param)

	if st.Len() > int(iparam) {
		return errors.New("field has more than " + param + " entries")
	}

	return nil

}

func customValidatorForEmail(v interface{}, param string) error {

	st := reflect.ValueOf(v)

	if !isValidEmail(st.String()) {
		return errors.New("field is not a valid email address")
	}

	return nil
}

func customValidatorForNoSpaces(v interface{}, param string) error {

	st := reflect.ValueOf(v)

	if strings.ContainsAny(st.String(), " ") {
		return errors.New("field should not have spaces it it")
	}

	return nil
}

func customValidatorForMaxFloat(v interface{}, param string) error {

	st := reflect.ValueOf(v)

	if st.Float() > math.MaxFloat64 {
		return errors.New("field has a value bigger than what the system can handle")
	}

	return nil
}

func customValidatorForAllowedCurrencies(v interface{}, param string) error {

	st := reflect.ValueOf(v)

	var currency []string

	//This list will add more values in the near future
	currency = append(currency, "USD")
	currency = append(currency, "CAD")
	currency = append(currency, "CDN")
	currency = append(currency, "EUR")
	currency = append(currency, "INR")
	currency = append(currency, "GBP")

	if !containsInArray(currency, st.String()) {
		return errors.New("field value " + st.String() + " not an acceptable currency")
	}

	return nil
}

func customValidatorForStarRating(v interface{}, param string) error {

	st := reflect.ValueOf(v)

	if st.Float() <= 0.0 || st.Float() > 5.0 {
		return errors.New("field needs to have value between 1.0 and 5.0 (inclusive)")
	}

	return nil
}

func customValidatorForSortOrder(v interface{}, param string) error {

	st := reflect.ValueOf(v)

	if !(st.Int() == -1 || st.Int() == 1) {
		return errors.New("field can only have either -1 (decending) or 1 (ascending) as a value")
	}

	return nil
}
