package main

import (
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/romana/rlog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//PATHSEPARATOR category path separator
const PATHSEPARATOR = ">"

func parseCategoryPath(path string, separator string) []string {

	return strings.Split(path, separator)

}

func deleteSKUFromTree(w http.ResponseWriter, r *http.Request, treeCollection string, path string, sku string) {

	rlog.Debug("deleteSKUFromTree() handle function invoked ...")

	catPath := parseCategoryPath(path, PATHSEPARATOR)
	pathLength := len(catPath)

	for i := 0; i < pathLength; i++ {

		node := getCategoryNode(w, r, catPath[i], treeCollection)

		if containsInArray(node.SKUs, sku) {

			if len(node.SKUs) == 1 && (node.Children == nil || len(node.Children) == 0) {

				node.SKUs = nil

				// Backtrack and keep on removing up until you hit a node with more than 1 children and/or SKUs in the node
				for node != nil && len(node.SKUs) == 0 && len(node.Children) == 0 {

					var names []string
					var delNodeName = node.Name
					deleteMongoDocument(ExternalDB, treeCollection, bson.M{"CategoryID": node.CategoryID})
					node = getCategoryNode(w, r, node.Parent, treeCollection)

					if node != nil {
						node.Children = removeElementsFromArray(node.Children, append(names, delNodeName))
						updateCategoryNode(w, r, node.CategoryID, treeCollection, node)
						node = getCategoryNode(w, r, node.Name, treeCollection)
					}
				}

			} else {

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

}

func insertIntoTree(w http.ResponseWriter, r *http.Request, treeCollection string, path string, sku string) bool {

	rlog.Debug("insertIntoTree() handle function invoked ...")

	catPath := parseCategoryPath(path, PATHSEPARATOR)
	pathLength := len(catPath)

	node := getCategoryNode(w, r, catPath[0], treeCollection)

	if node != nil && node.Parent != "" {
		respondWith(w, r, nil, "Root node in the category path is an existing child node in the tree", nil, http.StatusBadRequest, false)
		return false
	}

	for i := 0; i < pathLength; i++ {

		node := getCategoryNode(w, r, catPath[i], treeCollection)

		if node != nil && i == pathLength-1 {
			if !containsInArray(node.SKUs, sku) {
				node.SKUs = append(node.SKUs, sku)

				if node.Parent == "" && i-1 > 0 {
					node.Parent = catPath[i-1]
				}

				updateCategoryNode(w, r, node.CategoryID, treeCollection, node)
			}
		} else if node != nil && i < pathLength-1 {
			if !containsInArray(node.Children, catPath[i+1]) {
				node.Children = append(node.Children, catPath[i+1])

				if node.Parent == "" && i-1 > 0 {
					node.Parent = catPath[i-1]
				}

				updateCategoryNode(w, r, node.CategoryID, treeCollection, node)
			}
		} else if node == nil && i == pathLength-1 {

			var nCN CATEGORYTREENODE
			nCN.CategoryID = uuid.New().String()
			nCN.Name = catPath[i]
			nCN.SKUs = append(nCN.SKUs, sku)

			if i-1 >= 0 {
				nCN.Parent = catPath[i-1]
			}

			createCategoryNode(w, r, treeCollection, &nCN)

		} else if node == nil && i < pathLength-1 {

			var nCN CATEGORYTREENODE
			nCN.CategoryID = uuid.New().String()
			nCN.Name = catPath[i]
			nCN.Children = append(nCN.Children, catPath[i+1])

			if i-1 >= 0 {
				nCN.Parent = catPath[i-1]
			}

			createCategoryNode(w, r, treeCollection, &nCN)

		}

	}

	return true

}

func getCategoryNode(w http.ResponseWriter, r *http.Request, category string, collection string) *CATEGORYTREENODE {

	rlog.Debug("getCategoryNode() handle function invoked ...")

	var opts options.FindOptions

	results := findMongoDocument(ExternalDB, collection, bson.M{"Name": category}, &opts)

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

func getRootCategories(w http.ResponseWriter, r *http.Request, collection string) []string {

	rlog.Debug("getRootCategories() handle function invoked ...")

	var opts options.FindOptions

	results := findMongoDocument(ExternalDB, collection, bson.M{"Parent": ""}, &opts)

	var cats []string

	for _, result := range results {

		var node CATEGORYTREENODE

		mapDocument(w, r, &node, result)

		cats = append(cats, node.Name)

	}

	return cats

}

func pathExists(w http.ResponseWriter, r *http.Request, path string, collection string) bool {

	rlog.Debug("pathExists() handle function invoked ...")

	catPath := parseCategoryPath(path, PATHSEPARATOR)
	pathLength := len(catPath)

	if pathLength == 0 {
		rlog.Error("pathExists() path seems empty! ...")
		return false
	}

	for i := 0; i < pathLength; i++ {

		node := getCategoryNode(w, r, catPath[i], collection)

		if node == nil {
			return false
		}

	}

	return true
}

func getSKUsInTheCategoryPath(w http.ResponseWriter, r *http.Request, path string, collection string, onlyLeafSKUs bool) []string {

	rlog.Debug("getSKUsInTheCategoryPath() handle function invoked ...")

	catPath := parseCategoryPath(path, PATHSEPARATOR)
	pathLength := len(catPath)

	if pathLength == 0 {
		rlog.Error("getSKUsInTheCategoryPath() path seems empty! ...")
		return nil
	}

	var SKUs []string

	if !onlyLeafSKUs {

		for i := 0; i < pathLength; i++ {

			node := getCategoryNode(w, r, catPath[i], collection)

			if len(node.SKUs) != 0 {
				SKUs = append(SKUs, node.SKUs...)
			}
		}

	} else {

		node := getCategoryNode(w, r, catPath[pathLength-1], collection)
		SKUs = append(SKUs, node.SKUs...)

	}

	return SKUs

}
