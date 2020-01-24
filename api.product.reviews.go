package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/romana/rlog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getProductReviews(w http.ResponseWriter, r *http.Request) {

	rlog.Debug("getProductReviews() handle function invoked ...")

	if !pre(w, r) {
		return
	}

	var prr PRODREVIEWREQ

	if !mapInput(w, r, &prr) {
		return
	}

	if !validateProductReviewRequest(w, r, prr) {
		return
	}

	pth := strings.Split(r.URL.Path, "/")
	pgid := pth[len(pth)-1]

	dbcol := getAccessToken(r) + ProductReviewsExtension

	var opts options.FindOptions

	opts.SetSort(bson.M{prr.SortField: prr.Order})
	opts.SetSkip(prr.From)
	opts.SetLimit(prr.To)

	results := findMongoDocument(ExternalDB, dbcol, bson.M{"GroupID": pgid}, &opts)

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

	csx := getAccessToken(r)

	prdbcol := csx + ProductReviewsExtension
	pgdbcol := csx + ProductGroupExtension

	if !mapInput(w, r, &review) {
		return
	}

	review.ReviewID = uuid.New().String()

	if !validateProductReview(w, r, review) {
		return
	}

	review.Time = time.Now().UnixNano()

	var opts options.FindOptions

	results := findMongoDocument(ExternalDB, pgdbcol, bson.M{"GroupID": review.GroupID}, &opts)

	if len(results) != 1 {
		respondWith(w, r, nil, "Product Review Insertion Failed! Reason:"+ProductGroupNotFoundMessage, nil, http.StatusNotFound, false)
		return
	}

	insertMongoDocument(ExternalDB, prdbcol, review)

	var productGroup PRODUCTGROUP

	mapDocument(w, r, &productGroup, results[0])

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

	dbcol := getAccessToken(r) + ProductReviewsExtension

	pth := strings.Split(r.URL.Path, "/")
	rid := pth[len(pth)-1]

	var opts options.FindOptions

	results := findMongoDocument(ExternalDB, dbcol, bson.M{"ReviewID": rid}, &opts)

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

func deleteProductGroupReview(w http.ResponseWriter, r *http.Request) {

	rlog.Debug("deleteProductGroupReview() handle function invoked ...")

	if !pre(w, r) {
		return
	}

	dbcol := getAccessToken(r) + ProductReviewsExtension

	pth := strings.Split(r.URL.Path, "/")
	pgid := pth[len(pth)-1]

	var opts options.FindOptions

	results := findMongoDocument(ExternalDB, dbcol, bson.M{"GroupID": pgid}, &opts)

	if len(results) == 0 {
		respondWith(w, r, nil, "Reviews for Product group mentioned in request, Not Found ...", nil, http.StatusNotFound, false)
		return
	}

	if deleteMongoDocument(ExternalDB, dbcol, bson.M{"GroupID": pgid}) != 0 {

		respondWith(w, r, nil, "Reviews for product group deleted ...", nil, http.StatusOK, true)

	} else {

		respondWith(w, r, nil, "Reviews for product group not deleted ...", nil, http.StatusNotModified, false)

	}

}
