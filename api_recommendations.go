package main

import (
	"fmt"
	"net/http"

	"github.com/romana/rlog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getSimilarProducts(w http.ResponseWriter, r *http.Request) {

	rlog.Debug("getSimilarProducts() handle function invoked ...")

	if !pre(w, r) {
		return
	}

	var sirq SIRQ

	if !mapInput(w, r, &sirq) {
		return
	}

	csx := getAccessToken(r)
	ctcol := csx + CategoryTreeExtension
	pgcol := csx + ProductGroupExtension
	path := cleanCategoryPath(sirq.CategoryPath)

	if !pathExists(w, r, path, ExternalDB+csx, ctcol) {
		respondWith(w, r, nil, "Category path does not exit ...", nil, http.StatusBadRequest, false)
		return
	}

	SKUs := getSKUsInTheCategoryPath(w, r, path, ExternalDB+csx, ctcol, true)

	fmt.Printf("%+v", SKUs)

	var productG []PRODUCTGROUP

	var opts options.FindOptions

	var sx []bson.M

	for _, sku := range SKUs {
		sx = append(sx, bson.M{"Skus": sku})
	}

	results := findMongoDocument(ExternalDB+csx, pgcol, bson.M{"$or": sx}, &opts)

	if len(results) == 0 {
		respondWith(w, r, nil, ProductGroupNotFoundMessage, nil, http.StatusNotFound, false)
		return
	}

	for _, result := range results {

		var pg PRODUCTGROUP
		mapDocument(w, r, &pg, result)
		productG = append(productG, pg)

	}

	respondWith(w, r, nil, "Similar Products ...", productG, http.StatusOK, true)

}
