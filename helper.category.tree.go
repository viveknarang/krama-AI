package main

import (
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/romana/rlog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func parseCategoryPath(path string, separator string) []string {

	return strings.Split(path, separator)

}

func deleteSKUFromTree(w http.ResponseWriter, r *http.Request, treeCollection string, path string, sku string) {

	rlog.Debug("deleteSKUFromTree() handle function invoked ...")

	catPath := parseCategoryPath(path, ">")
	pathLength := len(catPath)

	for i := 0; i < pathLength; i++ {

		node := getCategoryNode(w, r, catPath[i], treeCollection)

		if containsInArray(node.SKUs, sku) {

			var fSKUs []string

			for _, v := range node.SKUs {
				if v != sku {
					fSKUs = append(fSKUs, v)
				}
			}

			node.SKUs = fSKUs
			updateCategoryNode(w, r, node.CategoryID, treeCollection, node)
		}

	}

}

func insertIntoTree(w http.ResponseWriter, r *http.Request, treeCollection string, path string, sku string) {

	rlog.Debug("insertIntoTree() handle function invoked ...")

	catPath := parseCategoryPath(path, ">")
	pathLength := len(catPath)

	for i := 0; i < pathLength; i++ {

		node := getCategoryNode(w, r, catPath[i], treeCollection)

		if node != nil && i == pathLength-1 {
			if !containsInArray(node.SKUs, sku) {
				node.SKUs = append(node.SKUs, sku)
				updateCategoryNode(w, r, node.CategoryID, treeCollection, node)
			}
		} else if node != nil && i < pathLength-1 {
			if !containsInArray(node.ChildCategory, catPath[i+1]) {
				node.ChildCategory = append(node.ChildCategory, catPath[i+1])
				updateCategoryNode(w, r, node.CategoryID, treeCollection, node)
			}
		} else if node == nil && i == pathLength-1 {

			var nCN CATEGORYTREENODE
			nCN.CategoryID = uuid.New().String()
			nCN.CategoryName = catPath[i]
			nCN.SKUs = append(nCN.SKUs, sku)

			if i-1 >= 0 {
				nCN.ParentCategory = catPath[i-1]
			}

			createCategoryNode(w, r, treeCollection, &nCN)

		} else if node == nil && i < pathLength-1 {

			var nCN CATEGORYTREENODE
			nCN.CategoryID = uuid.New().String()
			nCN.CategoryName = catPath[i]
			nCN.ChildCategory = append(nCN.ChildCategory, catPath[i+1])

			if i-1 >= 0 {
				nCN.ParentCategory = catPath[i-1]
			}

			createCategoryNode(w, r, treeCollection, &nCN)

		}

	}

}

func getCategoryNode(w http.ResponseWriter, r *http.Request, category string, collection string) *CATEGORYTREENODE {

	var opts options.FindOptions

	results := findMongoDocument(ExternalDB, collection, bson.M{"CategoryName": category}, &opts)

	if len(results) == 1 {

		var treeNode CATEGORYTREENODE

		mapDocument(w, r, &treeNode, results[0])

		return &treeNode
	}

	return nil

}

func updateCategoryNode(w http.ResponseWriter, r *http.Request, categoryID string, collection string, node *CATEGORYTREENODE) [2]int64 {

	return updateMongoDocument(ExternalDB, collection, bson.M{"CategoryID": categoryID}, bson.M{"$set": node})

}

func createCategoryNode(w http.ResponseWriter, r *http.Request, collection string, node *CATEGORYTREENODE) {

	insertMongoDocument(ExternalDB, collection, node)

}

func getRootCategories(w http.ResponseWriter, r *http.Request, collection string) *CATEGORYTREENODE {

	var opts options.FindOptions

	results := findMongoDocument(ExternalDB, collection, bson.M{"Parent": ""}, &opts)

	if len(results) == 1 {

		var treeNode CATEGORYTREENODE

		mapDocument(w, r, &treeNode, results[0])

		return &treeNode
	}

	return nil

}
