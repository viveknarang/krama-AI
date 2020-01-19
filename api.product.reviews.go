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

	csx := REDISCLIENT.Get(r.Header.Get("x-access-token")).Val()

	prdbcol := csx + ProductReviewsExtension
	pgdbcol := csx + ProductGroupExtension

	err := json.NewDecoder(r.Body).Decode(&review)

	if err != nil {
		respondWith(w, r, err, HTTPBadRequestMessage, nil, http.StatusBadRequest, false)
		return
	}

	if !validateProductReview(w, r, review) {
		return
	}

	review.Time = time.Now().UnixNano()
	review.ReviewID = uuid.New().String()

	insertMongoDocument(ExternalDB, prdbcol, review)

	results := findMongoDocument(ExternalDB, pgdbcol, bson.M{"GroupID": review.GroupID})

	if len(results) != 1 {
		respondWith(w, r, nil, "Product Review Insertion Failed! Reason:"+ProductGroupNotFoundMessage, nil, http.StatusNotFound, false)
		return
	}

	j, err0 := bson.MarshalExtJSON(results[0], false, false)

	if err0 != nil {
		respondWith(w, r, err0, HTTPInternalServerErrorMessage, nil, http.StatusInternalServerError, false)
		return
	}

	var productGroup PRODUCTGROUP

	err1 := json.Unmarshal([]byte(j), &productGroup)

	if err1 != nil {
		respondWith(w, r, err1, HTTPInternalServerErrorMessage, nil, http.StatusInternalServerError, false)
		return
	}

	newReviewCount := productGroup.CumulativeReviewCount + 1
	newReviewStars := (productGroup.CumulativeReviewStars + review.Stars) / float64(newReviewCount)

	updateResult := updateMongoDocument(ExternalDB, pgdbcol, bson.M{"GroupID": review.GroupID}, bson.M{"$set": bson.M{"CumulativeReviewStars": newReviewStars, "CumulativeReviewCount": newReviewCount}})

	if updateResult[0] == 1 && updateResult[1] == 1 {

		resetProductCacheKeys(nil, &productGroup)
		respondWith(w, r, nil, "Review Added and Cumulative review data updated in Product Group object ...", review, http.StatusCreated, true)

	} else if updateResult[0] == 1 && updateResult[1] == 0 {

		respondWith(w, r, nil, "Review Added and Cumulative review data not updated in Product Group object ...", review, http.StatusCreated, false)

	} else if updateResult[0] == 0 && updateResult[1] == 0 {

		respondWith(w, r, nil, "Review Added and Cumulative review data not updated in Product Group object ...", review, http.StatusCreated, false)

	}

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

	if len(results) == 0 {
		respondWith(w, r, nil, "Review Not Found ...", nil, http.StatusNotFound, false)
		return
	}

	if deleteMongoDocument(ExternalDB, dbcol, bson.M{"ReviewID": rid}) == 1 {

		respondWith(w, r, nil, "Review deleted ...", nil, http.StatusOK, true)

	} else {

		respondWith(w, r, nil, "Review not deleted ...", nil, http.StatusNotModified, false)

	}

}
