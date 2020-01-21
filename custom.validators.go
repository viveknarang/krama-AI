package main

import (
	"errors"
	"math"
	"reflect"
	"strconv"
	"strings"
)

func customValidatorForTypeArrayLengths(v interface{}, param string) error {

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

	var currency [6]string

	//This list will add more values in the near future
	currency[0] = "USD"
	currency[1] = "CAD"
	currency[2] = "CDN"
	currency[3] = "EUR"
	currency[4] = "INR"
	currency[5] = "GBP"

	if !contains(currency, st.String()) {
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
