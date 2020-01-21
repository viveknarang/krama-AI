package main

import (
	"net/http"

	"github.com/go-validator/validator"
)

func validateProductReviewRequest(w http.ResponseWriter, r *http.Request, reviewRq PRODREVIEWREQ) bool {

	validator.SetValidationFunc("isValidSortOrder", customValidatorForSortOrder)

	if errs := validator.Validate(reviewRq); errs != nil {

		respondWith(w, r, nil, "Error(s) found in the review data: "+errs.Error(), nil, http.StatusBadRequest, false)
		return false

	}

	return true

}
