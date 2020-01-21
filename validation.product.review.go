package main

import (
	"net/http"

	"github.com/go-validator/validator"
)

func validateProductReview(w http.ResponseWriter, r *http.Request, review PRODUCTREVIEW) bool {

	validator.SetValidationFunc("validStarRating", customValidatorForStarRating)

	if errs := validator.Validate(review); errs != nil {

		respondWith(w, r, nil, "Error(s) found in the review data: "+errs.Error(), nil, http.StatusBadRequest, false)
		return false

	}

	return true

}
