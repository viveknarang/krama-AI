package main

import (
	"net/http"

	"github.com/romana/rlog"
	"go.mongodb.org/mongo-driver/bson"
)

func getProductSkusInCategory(w http.ResponseWriter, r *http.Request) {

	rlog.Debug("getProductsInCategory() handle function invoked ...")

	if !pre(w, r) {
		return
	}

	csx := getAccessToken(r)
	ctcol := csx + CategoryTreeExtension

	var ctrq CATEGORYREQUEST

	if !mapInput(w, r, &ctrq) {
		return
	}

	path := cleanCategoryPath(ctrq.Path)

	if !pathExists(w, r, path, ExternalDB+csx, ctcol) {
		respondWith(w, r, nil, "Category path does not exit ...", nil, http.StatusBadRequest, false)
		return
	}

	SKUs := getSKUsInTheCategoryPath(w, r, path, ExternalDB+csx, ctcol, true)

	respondWith(w, r, nil, "Products in category path ...", bson.M{path: SKUs}, http.StatusOK, true)

}

func getProductsInCategory(w http.ResponseWriter, r *http.Request) {

	rlog.Debug("getProductsInCategory() handle function invoked ...")

	if !pre(w, r) {
		return
	}

	csx := getAccessToken(r)
	ctcol := csx + CategoryTreeExtension

	var ctrq CATEGORYREQUEST

	if !mapInput(w, r, &ctrq) {
		return
	}

	path := cleanCategoryPath(ctrq.Path)

	if !pathExists(w, r, path, ExternalDB+csx, ctcol) {
		respondWith(w, r, nil, "Category path does not exit ...", nil, http.StatusBadRequest, false)
		return
	}

	products := getProductsInTheCategoryPath(w, r, path, ExternalDB+csx, ctcol, true, csx)

	respondWith(w, r, nil, "Products in category path ...", bson.M{path: products}, http.StatusOK, true)

}

func getRootCategory(w http.ResponseWriter, r *http.Request) {

	rlog.Debug("getRootCategory() handle function invoked ...")

	if !pre(w, r) {
		return
	}

	csx := getAccessToken(r)
	ctcol := csx + CategoryTreeExtension

	cats := getRootCategories(w, r, ExternalDB+csx, ctcol)

	respondWith(w, r, nil, "Root categories ...", cats, http.StatusOK, true)

}

func getImmediateSubCategories(w http.ResponseWriter, r *http.Request) {

	rlog.Debug("getImmediateSubCategories() handle function invoked ...")

	if !pre(w, r) {
		return
	}

	csx := getAccessToken(r)
	ctcol := csx + CategoryTreeExtension

	var ctrq CATEGORYREQUEST

	if !mapInput(w, r, &ctrq) {
		return
	}

	catNode := getCategoryNode(w, r, ctrq.Category, ExternalDB+csx, ctcol)
	var childNodes []*CATEGORYTREENODE

	if catNode.Children == nil {

		respondWith(w, r, nil, "Category "+ctrq.Category+" does not have a sub category ...", nil, http.StatusNotFound, false)
		return

	}

	for _, child := range catNode.Children {
		childNodes = append(childNodes, getCategoryNode(w, r, child, ExternalDB+csx, ctcol))
	}

	respondWith(w, r, nil, "Immediate Sub categories ...", childNodes, http.StatusOK, true)

}

func getParentCategory(w http.ResponseWriter, r *http.Request) {

	rlog.Debug("getParentCategory() handle function invoked ...")

	if !pre(w, r) {
		return
	}

	csx := getAccessToken(r)
	ctcol := csx + CategoryTreeExtension

	var ctrq CATEGORYREQUEST

	if !mapInput(w, r, &ctrq) {
		return
	}

	catNode := getCategoryNode(w, r, ctrq.Category, ExternalDB+csx, ctcol)
	var parentNode *CATEGORYTREENODE

	if catNode.Parent == "" {

		respondWith(w, r, nil, "Category "+ctrq.Category+" does not have a parent ...", nil, http.StatusNotFound, false)
		return

	} else {

		parentNode = getCategoryNode(w, r, catNode.Parent, ExternalDB+csx, ctcol)

	}

	respondWith(w, r, nil, "Category parent ...", parentNode, http.StatusOK, true)

}
