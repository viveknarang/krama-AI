package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/romana/rlog"
	"go.mongodb.org/mongo-driver/bson"
)

func getProductReviews(w http.ResponseWriter, r *http.Request) {

	rlog.Debug("getProductReviews() handle function invoked ...")

	if !pre(w, r) {
		return
	}

	pth := strings.Split(r.URL.Path, "/")
	pgid := pth[len(pth)-1]

	dbcol := REDISCLIENT.Get(r.Header.Get("x-access-token")).Val() + ProductReviewsExtension

	results := findMongoDocument(ExternalDB, dbcol, bson.M{"GroupID": pgid})

	if len(results) == 0 {
		respondWith(w, r, nil, "Reviews Not found ...", nil, http.StatusNotFound, false)
		return
	}

	var resultArr []interface{}

	for _, rslt := range results {

		j, err0 := bson.MarshalExtJSON(rslt, false, false)

		if err0 != nil {
			respondWith(w, r, err0, HTTPInternalServerErrorMessage, nil, http.StatusInternalServerError, false)
			return
		}

		var result PRODUCTREVIEW

		err1 := json.Unmarshal([]byte(j), &result)

		if err1 != nil {
			respondWith(w, r, err1, HTTPInternalServerErrorMessage, nil, http.StatusInternalServerError, false)
			return
		}

		resultArr = append(resultArr, result)

	}

	respondWith(w, r, nil, CustomersFoundMessage, resultArr, http.StatusOK, false)

}

func postProductReview(w http.ResponseWriter, r *http.Request) {

	rlog.Debug("postProductReview() handle function invoked ...")

	if !pre(w, r) {
		return
	}

	var review PRODUCTREVIEW

	dbcol := REDISCLIENT.Get(r.Header.Get("x-access-token")).Val() + ProductReviewsExtension

	err := json.NewDecoder(r.Body).Decode(&review)

	if err != nil {
		respondWith(w, r, err, HTTPBadRequestMessage, nil, http.StatusBadRequest, false)
		return
	}

	//if !validateReview() {
	//	return
	//}

	review.Time = time.Now().UnixNano()
	review.ReviewID = uuid.New().String()

	insertMongoDocument(ExternalDB, dbcol, review)

	respondWith(w, r, nil, "Review Added ...", review, http.StatusCreated, true)

}

func deleteProductReview(w http.ResponseWriter, r *http.Request) {

	rlog.Debug("deleteProductReview() handle function invoked ...")

	if !pre(w, r) {
		return
	}

	dbcol := REDISCLIENT.Get(r.Header.Get("x-access-token")).Val() + ProductReviewsExtension

	pth := strings.Split(r.URL.Path, "/")
	rid := pth[len(pth)-1]

	results := findMongoDocument(ExternalDB, dbcol, bson.M{"ReviewID": rid})

	if len(results) != 1 {
		respondWith(w, r, nil, "Review Not Found ...", nil, http.StatusNotFound, false)
		return
	}

	if deleteMongoDocument(ExternalDB, dbcol, bson.M{"ReviewID": rid}) == 1 {

		respondWith(w, r, nil, "Review deleted ...", nil, http.StatusOK, true)

	} else {

		respondWith(w, r, nil, "Review not deleted ...", nil, http.StatusNotModified, false)

	}

}
