package main

import (
	"net/http"
	"strings"

	"github.com/go-validator/validator"
)

func validateProduct(w http.ResponseWriter, r *http.Request, product PRODUCT) bool {

	if errs := validator.Validate(product); errs != nil {

		respondWith(w, r, nil, "Error(s) found in the product data: "+errs.Error(), nil, http.StatusBadRequest, false)
		return false

	}

	for _, img := range product.Images {

		if !isValidURL(img) {
			respondWith(w, r, nil, "Image URL: "+img+" is not a valid URL", nil, http.StatusBadRequest, false)
			return false
		}

	}

	if len(product.Attributes) > 0 {

		for key, value := range product.Attributes {
			if strings.Contains(typeof(key), "interface") || strings.Contains(typeof(value), "interface") {
				respondWith(w, r, nil, "Attribute field keys or values cannot be complex object. They need to be simple types like int, float or boolean etc ...", nil, http.StatusBadRequest, false)
				return false
			}
			if !isValidAttributeKey(key) {
				respondWith(w, r, nil, "Attribute field key: "+key+" is not following attribute naming rules. Please check API documentation.", nil, http.StatusBadRequest, false)
				return false
			}
		}

	}

	return true

}
