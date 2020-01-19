package main

import (
	"net/http"
)

func validateProductReviewRequest(w http.ResponseWriter, r *http.Request, req PRODREVIEWREQ) bool {

	if len(req.SortField) == 0 || len(req.SortField) > 100 {

		respondWith(w, r, nil, "SortField field cannot be empty or greater than 100 characters long", nil, http.StatusBadRequest, false)
		return false

	}

	if !(req.Order == -1 || req.Order == 1) {

		respondWith(w, r, nil, "Order field can only have either -1 (decending) or 1 (ascending) as a value", nil, http.StatusBadRequest, false)
		return false

	}

	return true

}
