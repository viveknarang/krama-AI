package main

import (
	"encoding/json"
	"math"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

func isValidJSON(s string) bool {

	var js interface{}
	return json.Unmarshal([]byte(s), &js) == nil

}

func isValidURL(toTest string) bool {

	_, err := url.ParseRequestURI(toTest)

	return !(err != nil)

}

func typeof(value interface{}) string {

	return reflect.TypeOf(value).String()

}

func pre(w http.ResponseWriter, r *http.Request) bool {

	if (r.Method == http.MethodPost || r.Method == http.MethodPut) && r.Header.Get("Content-Type") != "application/json" {

		respondWith(w, r, nil, MissingContentType, nil, http.StatusBadRequest)
		return false

	}

	if r.Header.Get("x-access-token") == "" {

		respondWith(w, r, nil, MissingAccessToken, nil, http.StatusBadRequest)
		return false

	}

	if !authenticate(r.Header.Get("x-access-token")) {

		respondWith(w, r, nil, InvalidSessionMessage, nil, http.StatusUnauthorized)
		return false

	}

	return true

}

func validateProduct(w http.ResponseWriter, r *http.Request, product PRODUCT) bool {

	if len(product.Name) == 0 || len(product.Name) > 100 {

		respondWith(w, r, nil, "Product Name field cannot be empty or greater than 100 characters long", nil, http.StatusBadRequest)
		return false

	}

	if len(product.Description) > 10240 {

		respondWith(w, r, nil, "Description field cannot be greater than 10240 characters long", nil, http.StatusBadRequest)
		return false

	}

	if len(product.Sku) == 0 || len(product.Sku) > 50 || strings.ContainsAny(product.Sku, " ") {

		respondWith(w, r, nil, "Sku field cannot contain spaces, be empty, or greater than 50 characters long", nil, http.StatusBadRequest)
		return false

	}

	if len(product.GroupID) == 0 || len(product.GroupID) > 50 || strings.ContainsAny(product.GroupID, " ") {

		respondWith(w, r, nil, "Sku field cannot contain spaces, be empty, or greater than 50 characters long", nil, http.StatusBadRequest)
		return false

	}

	if product.Quantity < 0 {

		respondWith(w, r, nil, "Quantity field cannot have a negative value", nil, http.StatusBadRequest)
		return false

	}

	if len(product.Category) == 0 || len(product.Category) > 100 {

		respondWith(w, r, nil, "Category field for any product cannot be empty  or contain more than 100 categories", nil, http.StatusBadRequest)
		return false

	}

	if len(product.Images) == 0 || len(product.Images) > 100 {

		respondWith(w, r, nil, "Images field for any product cannot be empty or contain more than 100 image URLs", nil, http.StatusBadRequest)
		return false

	}

	for _, img := range product.Images {
		if !isValidURL(img) {
			respondWith(w, r, nil, "Image URL "+img+" is not a valid URL", nil, http.StatusBadRequest)
			return false
		}
	}

	if product.PromotionPrice < 0.0 || product.PromotionPrice > math.MaxFloat64 {

		respondWith(w, r, nil, "PromotionPrice field cannot have a negative value or greater than the maximum possible value", nil, http.StatusBadRequest)
		return false

	}

	if product.RegularPrice < 0.0 || product.RegularPrice > math.MaxFloat64 {

		respondWith(w, r, nil, "RegularPrice field cannot have a negative value or greater than the maximum possible value", nil, http.StatusBadRequest)
		return false

	}

	if len(product.SearchKeywords) == 0 || len(product.SearchKeywords) > 100 {

		respondWith(w, r, nil, "SearchKeywords field cannot be empty or contain more than 100 search keywords", nil, http.StatusBadRequest)
		return false

	}

	if len(product.Attributes) > 0 {

		for key, value := range product.Attributes {
			if strings.Contains(typeof(key), "interface") || strings.Contains(typeof(value), "interface") {
				respondWith(w, r, nil, "Attribute field keys or values cannot be complex object. They need to be simple types like int, float or boolean etc ...", nil, http.StatusBadRequest)
				return false
			}
		}

	}

	if len(product.Attributes) > 500 {

		respondWith(w, r, nil, "A product entity cannot contain more than 500 additional attributes", nil, http.StatusBadRequest)
		return false

	}

	return true

}
