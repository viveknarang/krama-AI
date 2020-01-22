package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/go-redis/redis"
	"github.com/romana/rlog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getProduct(w http.ResponseWriter, r *http.Request) {

	rlog.Debug("getProduct() handle function invoked ...")

	if !pre(w, r) {
		return
	}

	var jx []byte
	var product PRODUCT

	redisC := REDISCLIENT.Get(r.URL.Path)
	csx := getAccessToken(r)

	if redisC.Err() != redis.Nil {

		jx = []byte(redisC.Val())
		mapBytes(w, r, &product, jx)

	} else {

		pth := strings.Split(r.URL.Path, "/")
		sku := pth[len(pth)-1]

		dbcol := csx + ProductExtension

		var opts options.FindOptions

		results := findMongoDocument(ExternalDB, dbcol, bson.M{"Sku": sku}, &opts)

		if len(results) != 1 {
			respondWith(w, r, nil, ProductNotFoundMessage, nil, http.StatusNotFound, false)
			return
		}

		mapDocument(w, r, &product, results[0])

		jx = mapToBytes(w, r, results[0])

		REDISCLIENT.Set(r.URL.Path, jx, 0)

	}

	picol := csx + ProductInventoryExtension
	var opts options.FindOptions

	results := findMongoDocument(ExternalDB, picol, bson.M{"Sku": product.Sku}, &opts)

	if len(results) != 1 {
		respondWith(w, r, nil, "Inventory Record Not found ...", nil, http.StatusNotFound, false)
		return
	}

	var inventory INVENTORY

	mapDocument(w, r, &inventory, results[0])

	product.Quantity = inventory.Quantity

	respondWith(w, r, nil, ProductFoundMessage, product, http.StatusOK, false)

}

func postProduct(w http.ResponseWriter, r *http.Request) {

	rlog.Debug("postProduct() handle function invoked ...")

	if !pre(w, r) {
		return
	}

	csx := getAccessToken(r)
	dbcol := csx + ProductExtension
	picol := csx + ProductInventoryExtension

	var p PRODUCT

	mapInput(w, r, &p)

	var opts options.FindOptions

	results := findMongoDocument(ExternalDB, dbcol, bson.M{"Sku": p.Sku}, &opts)

	if len(results) != 0 {
		respondWith(w, r, nil, ProductAlreadyExistsMessage, nil, http.StatusConflict, false)
		return
	}

	if !validateProduct(w, r, p) {
		return
	}

	groomProductData(&p)

	p.Updated = time.Now().UnixNano()

	insertMongoDocument(ExternalDB, dbcol, p)

	for _, cat := range p.Category {
		insertIntoTree(w, r, csx+CategoryTreeExtension, cat, p.Sku)
	}

	var productInventoryRecord INVENTORY
	productInventoryRecord.Sku = p.Sku
	productInventoryRecord.Quantity = p.Quantity
	productInventoryRecord.Updated = time.Now().UnixNano()
	insertMongoDocument(ExternalDB, picol, productInventoryRecord)

	if syncProductGroup(w, r, p) {

		resetProductCacheKeys(&p, nil)

		respondWith(w, r, nil, ProductAddedMessage, p, http.StatusCreated, true)

	} else {

		respondWith(w, r, nil, ProductNotAddedMessage, p, http.StatusNotModified, false)

	}

}

func putProduct(w http.ResponseWriter, r *http.Request) {

	rlog.Debug("putProduct() handle function invoked ...")

	if !pre(w, r) {
		return
	}

	var p PRODUCT

	mapInput(w, r, &p)

	if !validateProduct(w, r, p) {
		return
	}

	groomProductData(&p)

	dbcol := getAccessToken(r) + ProductExtension

	result := updateMongoDocument(ExternalDB, dbcol, bson.M{"Sku": p.Sku}, bson.M{"$set": p})

	if result[0] == 1 && result[1] == 1 {

		if syncProductGroup(w, r, p) {

			resetProductCacheKeys(&p, nil)
			respondWith(w, r, nil, ProductUpdatedMessage, p, http.StatusAccepted, true)

		} else {

			respondWith(w, r, nil, ProductNotUpdatedMessage, p, http.StatusNotModified, false)

		}

	} else if result[0] == 1 && result[1] == 0 {

		respondWith(w, r, nil, ProductNotUpdatedMessage, nil, http.StatusNotModified, false)

	} else if result[0] == 0 && result[1] == 0 {

		respondWith(w, r, nil, ProductNotFoundMessage, nil, http.StatusNotModified, false)

	}

}

func deleteProduct(w http.ResponseWriter, r *http.Request) {

	rlog.Debug("deleteProduct() handle function invoked ...")

	if !pre(w, r) {
		return
	}

	csx := getAccessToken(r)
	dbcol := csx + ProductExtension
	picol := csx + ProductInventoryExtension
	ctcol := csx + CategoryTreeExtension

	pth := strings.Split(r.URL.Path, "/")
	sku := pth[len(pth)-1]

	var opts options.FindOptions

	results := findMongoDocument(ExternalDB, dbcol, bson.M{"Sku": sku}, &opts)

	if len(results) != 1 {
		respondWith(w, r, nil, ProductNotFoundMessage, nil, http.StatusNotFound, false)
		return
	}

	var product PRODUCT

	mapDocument(w, r, &product, results[0])

	if deleteMongoDocument(ExternalDB, dbcol, bson.M{"Sku": sku}) == 1 {

		deleteMongoDocument(ExternalDB, picol, bson.M{"Sku": sku})

		for _, cat := range product.Category {
			deleteSKUFromTree(w, r, ctcol, cat, product.Sku)
		}

		if syncProductGroup(w, r, product) {

			resetProductCacheKeys(&product, nil)
			respondWith(w, r, nil, ProductDeletedMessage, nil, http.StatusOK, true)

		} else {

			respondWith(w, r, nil, ProductNotDeletedMessage, nil, http.StatusOK, true)

		}

	} else {

		respondWith(w, r, nil, ProductNotFoundMessage, nil, http.StatusNotModified, false)

	}

}

func getProductGroup(w http.ResponseWriter, r *http.Request) {

	rlog.Debug("getProductGroup() handle function invoked ...")

	if !pre(w, r) {
		return
	}

	var jx []byte
	var productGroup PRODUCTGROUP

	redisC := REDISCLIENT.Get(r.URL.Path)

	if redisC.Err() != redis.Nil {

		jx = []byte(redisC.Val())
		mapBytes(w, r, &productGroup, jx)

	} else {

		dbcol := getAccessToken(r) + ProductGroupExtension

		pth := strings.Split(r.URL.Path, "/")
		pgid := pth[len(pth)-1]

		var opts options.FindOptions

		results := findMongoDocument(ExternalDB, dbcol, bson.M{"GroupID": pgid}, &opts)

		if len(results) != 1 {
			respondWith(w, r, nil, ProductGroupNotFoundMessage, nil, http.StatusNotFound, false)
			return
		}

		mapDocument(w, r, &productGroup, results[0])

		jx = mapToBytes(w, r, results[0])

		REDISCLIENT.Set(r.URL.Path, jx, 0)

	}

	respondWith(w, r, nil, ProductGroupFoundMessage, productGroup, http.StatusOK, true)

}

func deleteProductGroup(w http.ResponseWriter, r *http.Request) {

	rlog.Debug("deleteProductGroup() handle function invoked ...")

	if !pre(w, r) {
		return
	}

	cidb := getAccessToken(r)

	pgcol := cidb + ProductGroupExtension
	pcol := cidb + ProductExtension
	pgindex := cidb + ProductGroupExtension + SearchIndexExtension
	picol := cidb + ProductInventoryExtension
	ctcol := cidb + CategoryTreeExtension

	pth := strings.Split(r.URL.Path, "/")
	pgid := pth[len(pth)-1]

	var opts options.FindOptions

	results := findMongoDocument(ExternalDB, pgcol, bson.M{"GroupID": pgid}, &opts)

	if len(results) != 1 {
		respondWith(w, r, nil, ProductGroupNotFoundMessage, nil, http.StatusNotFound, false)
		return
	}

	var productGroup PRODUCTGROUP

	mapDocument(w, r, &productGroup, results[0])

	for _, product := range productGroup.Products {

		deleteMongoDocument(ExternalDB, pcol, bson.M{"Sku": product.Sku})
		deleteMongoDocument(ExternalDB, picol, bson.M{"Sku": product.Sku})

		for _, cat := range product.Category {
			deleteSKUFromTree(w, r, ctcol, cat, product.Sku)
		}

	}

	if deleteMongoDocument(ExternalDB, pgcol, bson.M{"GroupID": pgid}) == 1 {

		resetProductCacheKeys(nil, &productGroup)

		deleteESDocumentByID(pgindex, pgid)

		respondWith(w, r, nil, ProductGroupDeletedMessage, nil, http.StatusOK, true)

	} else {
		respondWith(w, r, nil, ProductGroupNotFoundMessage, nil, http.StatusNotModified, false)
	}

}

func updateProductsPrice(w http.ResponseWriter, r *http.Request) {

	rlog.Debug("updateProductsPrice() handle function invoked ...")

	if !pre(w, r) {
		return
	}

	var prices PRICEUPDATEREQUEST

	mapInput(w, r, &prices)

	for sku, price := range prices.Prices {
		if price.RegularPrice < 0 || price.PromotionPrice < 0 {
			respondWith(w, r, nil, "Price for sku: "+sku+" is negative. Prices cannot be negative ...", nil, http.StatusBadRequest, false)
			return
		}
	}

	dbcol := getAccessToken(r) + ProductExtension

	var priceUpdated []string
	var priceNotUpdated []string
	var priceNotFound []string

	for sku, price := range prices.Prices {

		result := updateMongoDocument(ExternalDB, dbcol, bson.M{"Sku": sku}, bson.M{"$set": bson.M{"RegularPrice": price.RegularPrice, "PromotionPrice": price.PromotionPrice}})

		if result[0] == 1 && result[1] == 1 {
			priceUpdated = append(priceUpdated, sku)
		} else if result[0] == 1 && result[1] == 0 {
			priceNotUpdated = append(priceNotUpdated, sku)
		} else if result[0] == 0 && result[1] == 0 {
			priceNotFound = append(priceNotFound, sku)
		}

	}

	if syncProductGroupFromProducts(w, r, priceUpdated, true) {

		respondWith(w, r, nil, "Prices Updated ...", bson.M{"Products Updated": priceUpdated, "Products Not Updated": priceNotUpdated, "Products Not Found": priceNotFound}, http.StatusOK, true)

	} else {

		respondWith(w, r, nil, "Prices Not updated ...", nil, http.StatusNotModified, false)

	}

}

func updateProductsInventory(w http.ResponseWriter, r *http.Request) {

	rlog.Debug("updateProductsInventory() handle function invoked ...")

	if !pre(w, r) {
		return
	}

	var quantities INVENTORYUPDATEREQUEST

	mapInput(w, r, &quantities)

	for sku, quantity := range quantities.Quantity {
		if quantity < 0 {
			respondWith(w, r, nil, "Inventory for sku: "+sku+" is negative. Quantity field cannot be negative ...", nil, http.StatusBadRequest, false)
			return
		}
	}

	picol := getAccessToken(r) + ProductInventoryExtension

	var quantityUpdated []string
	var quantityNotUpdated []string
	var quantityNotFound []string

	for sku, quantity := range quantities.Quantity {

		result := updateInventory(w, r, picol, "INCR", sku, quantity, true)

		if result[0] == 1 && result[1] == 1 {
			quantityUpdated = append(quantityUpdated, sku)
		} else if result[0] == 1 && result[1] == 0 {
			quantityNotUpdated = append(quantityNotUpdated, sku)
		} else if result[0] == 0 && result[1] == 0 {
			quantityNotFound = append(quantityNotFound, sku)
		}

	}

	if syncProductGroupFromProducts(w, r, quantityUpdated, false) {

		respondWith(w, r, nil, "Inventory Updated ...", bson.M{"Products Updated": quantityUpdated, "Products Not Updated": quantityNotUpdated, "Products Not Found": quantityNotFound}, http.StatusOK, true)

	} else {

		respondWith(w, r, nil, "Inventory Not updated ...", nil, http.StatusNotModified, false)

	}

}
