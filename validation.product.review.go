package main

import (
	"net/http"
)

func validateProductReview(w http.ResponseWriter, r *http.Request, review PRODUCTREVIEW) bool {

	if len(review.CustomerID) > 100 {

		respondWith(w, r, nil, "Customer ID field cannot be greater than 100 characters long", nil, http.StatusBadRequest, false)
		return false

	}

	if len(review.Description) > 10240 {

		respondWith(w, r, nil, "Description field cannot be greater than 10240 characters long", nil, http.StatusBadRequest, false)
		return false

	}

	if len(review.GroupID) == 0 || len(review.GroupID) > 100 {

		respondWith(w, r, nil, "Group ID field cannot be empty or greater than 100 characters long", nil, http.StatusBadRequest, false)
		return false

	}

	if len(review.ReviewID) > 100 {

		respondWith(w, r, nil, "Review ID field cannot be greater than 100 characters long", nil, http.StatusBadRequest, false)
		return false

	}

	if review.Stars < 1 || review.Stars > 5 {

		respondWith(w, r, nil, "Review stars field needs to be between 1 and 5 (inclusive)", nil, http.StatusBadRequest, false)
		return false

	}

	return true

}
