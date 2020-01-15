package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/bson"
)

func getProduct(w http.ResponseWriter, r *http.Request) {

	if !pre(w, r) {
		return
	}

	var jx []byte

	redisC := REDISCLIENT.Get(r.URL.Path)

	if redisC.Err() != redis.Nil {

		jx = []byte(redisC.Val())

	} else {

		pth := strings.Split(r.URL.Path, "/")
		sku := pth[len(pth)-1]

		dbcol := REDISCLIENT.Get(r.Header.Get("x-access-token")).Val() + ProductExtension

		results := findMongoDocument(ExternalDB, dbcol, bson.M{"sku": sku})

		if len(results) != 1 {
			respondWith(w, r, nil, ProductNotFoundMessage, nil, http.StatusNotFound, false)
			return
		}

		j, err0 := bson.MarshalExtJSON(results[0], false, false)

		if err0 != nil {
			respondWith(w, r, err0, HTTPInternalServerErrorMessage, nil, http.StatusInternalServerError, false)
			return
		}

		jx = j

		REDISCLIENT.Set(r.URL.Path, j, 0)

	}

	var product PRODUCT

	err1 := json.Unmarshal([]byte(jx), &product)

	if err1 != nil {
		respondWith(w, r, err1, HTTPInternalServerErrorMessage, nil, http.StatusInternalServerError, false)
		return
	}

	respondWith(w, r, nil, ProductFoundMessage, product, http.StatusOK, false)

}

func postProduct(w http.ResponseWriter, r *http.Request) {

	if !pre(w, r) {
		return
	}

	var p PRODUCT

	err := json.NewDecoder(r.Body).Decode(&p)

	if err != nil {
		respondWith(w, r, err, HTTPBadRequestMessage, nil, http.StatusBadRequest, false)
		return
	}

	if !validateProduct(w, r, p) {
		return
	}

	groomProductData(&p)

	p.Updated = time.Now().UnixNano()

	dbcol := REDISCLIENT.Get(r.Header.Get("x-access-token")).Val() + ProductExtension

	insertMongoDocument(ExternalDB, dbcol, p)

	if syncProductGroup(w, r, p) {

		resetProductCacheKeys(&p, nil)

		respondWith(w, r, nil, ProductAddedMessage, p, http.StatusCreated, true)

	} else {

		respondWith(w, r, nil, ProductNotAddedMessage, p, http.StatusNotModified, false)

	}

}

func putProduct(w http.ResponseWriter, r *http.Request) {

	if !pre(w, r) {
		return
	}

	var p PRODUCT

	err := json.NewDecoder(r.Body).Decode(&p)

	if err != nil {
		respondWith(w, r, err, HTTPBadRequestMessage, nil, http.StatusBadRequest, false)
		return
	}

	if !validateProduct(w, r, p) {
		return
	}

	groomProductData(&p)

	dbcol := REDISCLIENT.Get(r.Header.Get("x-access-token")).Val() + ProductExtension

	result := updateMongoDocument(ExternalDB, dbcol, bson.M{"sku": p.Sku}, bson.M{"$set": p})

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

	if !pre(w, r) {
		return
	}

	dbcol := REDISCLIENT.Get(r.Header.Get("x-access-token")).Val() + ProductExtension

	pth := strings.Split(r.URL.Path, "/")
	sku := pth[len(pth)-1]

	results := findMongoDocument(ExternalDB, dbcol, bson.M{"sku": sku})

	if len(results) != 1 {
		respondWith(w, r, nil, ProductNotFoundMessage, nil, http.StatusNotFound, false)
		return
	}

	j, err0 := bson.MarshalExtJSON(results[0], false, false)

	if err0 != nil {
		respondWith(w, r, err0, HTTPInternalServerErrorMessage, nil, http.StatusInternalServerError, false)
		return
	}

	var product PRODUCT

	err1 := json.Unmarshal([]byte(j), &product)

	if err1 != nil {
		respondWith(w, r, err1, HTTPInternalServerErrorMessage, nil, http.StatusInternalServerError, false)
		return
	}

	if deleteMongoDocument(ExternalDB, dbcol, bson.M{"sku": sku}) == 1 {

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

	if !pre(w, r) {
		return
	}

	var jx []byte

	redisC := REDISCLIENT.Get(r.URL.Path)

	if redisC.Err() != redis.Nil {

		jx = []byte(redisC.Val())

	} else {

		dbcol := REDISCLIENT.Get(r.Header.Get("x-access-token")).Val() + ProductGroupExtension

		pth := strings.Split(r.URL.Path, "/")
		pgid := pth[len(pth)-1]

		results := findMongoDocument(ExternalDB, dbcol, bson.M{"groupid": pgid})

		if len(results) != 1 {
			respondWith(w, r, nil, ProductGroupNotFoundMessage, nil, http.StatusNotFound, false)
			return
		}

		j, err0 := bson.MarshalExtJSON(results[0], false, false)

		if err0 != nil {
			respondWith(w, r, err0, HTTPInternalServerErrorMessage, nil, http.StatusInternalServerError, false)
			return
		}

		jx = j

		REDISCLIENT.Set(r.URL.Path, j, 0)

	}

	var productGroup PRODUCTGROUP

	err1 := json.Unmarshal([]byte(jx), &productGroup)

	if err1 != nil {
		respondWith(w, r, err1, HTTPInternalServerErrorMessage, nil, http.StatusInternalServerError, false)
		return
	}

	respondWith(w, r, nil, ProductGroupFoundMessage, productGroup, http.StatusOK, true)

}

func deleteProductGroup(w http.ResponseWriter, r *http.Request) {

	if !pre(w, r) {
		return
	}

	cidb := REDISCLIENT.Get(r.Header.Get("x-access-token")).Val()

	pgcol := cidb + ProductGroupExtension
	pcol := cidb + ProductExtension
	pgindex := cidb + ProductGroupExtension + SearchIndexExtension

	pth := strings.Split(r.URL.Path, "/")
	pgid := pth[len(pth)-1]

	results := findMongoDocument(ExternalDB, pgcol, bson.M{"groupid": pgid})

	if len(results) != 1 {
		respondWith(w, r, nil, ProductGroupNotFoundMessage, nil, http.StatusNotFound, false)
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

	for _, sku := range productGroup.Skus {
		deleteMongoDocument(ExternalDB, pcol, bson.M{"sku": sku})
	}

	if deleteMongoDocument(ExternalDB, pgcol, bson.M{"groupid": pgid}) == 1 {

		resetProductCacheKeys(nil, &productGroup)

		deleteESDocumentByID(pgindex, pgid)

		respondWith(w, r, nil, ProductGroupDeletedMessage, nil, http.StatusOK, true)

	} else {
		respondWith(w, r, nil, ProductGroupNotFoundMessage, nil, http.StatusNotModified, false)
	}

}

func updateProductsPrice(w http.ResponseWriter, r *http.Request) {

	if !pre(w, r) {
		return
	}

	var prices PRICEUPDATEREQUEST

	err := json.NewDecoder(r.Body).Decode(&prices)

	if err != nil {
		respondWith(w, r, err, HTTPBadRequestMessage, nil, http.StatusBadRequest, false)
		return
	}

	for sku, price := range prices.Prices {
		if price.RegularPrice < 0 || price.PromotionPrice < 0 {
			respondWith(w, r, err, "Price for sku: "+sku+" is negative. Prices cannot be negative ...", nil, http.StatusBadRequest, false)
			return
		}
	}

	dbcol := REDISCLIENT.Get(r.Header.Get("x-access-token")).Val() + ProductExtension

	var priceUpdated []string
	var priceNotUpdated []string
	var priceNotFound []string

	for sku, price := range prices.Prices {

		result := updateMongoDocument(ExternalDB, dbcol, bson.M{"sku": sku}, bson.M{"$set": bson.M{"regularprice": price.RegularPrice, "promotionprice": price.PromotionPrice}})

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

	if !pre(w, r) {
		return
	}

	var quantities INVENTORYUPDATEREQUEST

	err := json.NewDecoder(r.Body).Decode(&quantities)

	if err != nil {
		respondWith(w, r, err, HTTPBadRequestMessage, nil, http.StatusBadRequest, false)
		return
	}

	for sku, quantity := range quantities.Quantity {
		if quantity < 0 {
			respondWith(w, r, err, "Inventory for sku: "+sku+" is negative. Quantity field cannot be negative ...", nil, http.StatusBadRequest, false)
			return
		}
	}

	dbcol := REDISCLIENT.Get(r.Header.Get("x-access-token")).Val() + ProductExtension

	var quantityUpdated []string
	var quantityNotUpdated []string
	var quantityNotFound []string

	for sku, quantity := range quantities.Quantity {

		result := updateMongoDocument(ExternalDB, dbcol, bson.M{"sku": sku}, bson.M{"$set": bson.M{"quantity": quantity}})

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
