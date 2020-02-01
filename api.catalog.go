package main

import (
	"net/http"
	"strconv"
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

		results := findMongoDocument(ExternalDB+csx, dbcol, bson.M{"Sku": sku}, &opts)

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

	results := findMongoDocument(ExternalDB+csx, picol, bson.M{"Sku": product.Sku}, &opts)

	if len(results) != 1 {
		respondWith(w, r, nil, "Inventory Record Not found ...", nil, http.StatusNotFound, false)
		return
	}

	var inventory INVENTORY

	mapDocument(w, r, &inventory, results[0])

	product.Quantity = inventory.Quantity

	respondWith(w, r, nil, ProductFoundMessage, product, http.StatusOK, false)

}

func getProducts(w http.ResponseWriter, r *http.Request) {

	rlog.Debug("getProducts() handle function invoked ...")

	if !pre(w, r) {
		return
	}

	var prq PRQ

	if !mapInput(w, r, &prq) {
		return
	}

	var products []PRODUCT

	csx := getAccessToken(r)

	dbcol := csx + ProductExtension

	var opts options.FindOptions

	var sx []bson.M

	for _, sku := range prq.Skus {
		sx = append(sx, bson.M{"Sku": sku})
	}

	results := findMongoDocument(ExternalDB+csx, dbcol, bson.M{"$or": sx}, &opts)

	if len(results) == 0 {
		respondWith(w, r, nil, ProductsNotFoundMessage, nil, http.StatusNotFound, false)
		return
	}

	for _, result := range results {

		var p PRODUCT
		mapDocument(w, r, &p, result)
		products = append(products, p)

	}

	respondWith(w, r, nil, ProductsFoundMessage, products, http.StatusOK, false)

}

func getProductGroups(w http.ResponseWriter, r *http.Request) {

	rlog.Debug("getProductGroups() handle function invoked ...")

	if !pre(w, r) {
		return
	}

	var pgrq PGRQ

	if !mapInput(w, r, &pgrq) {
		return
	}

	var productG []PRODUCTGROUP

	csx := getAccessToken(r)

	dbcol := csx + ProductGroupExtension

	var opts options.FindOptions

	var sx []bson.M

	for _, sku := range pgrq.Skus {
		sx = append(sx, bson.M{"Skus": sku})
	}

	results := findMongoDocument(ExternalDB+csx, dbcol, bson.M{"$or": sx}, &opts)

	if len(results) == 0 {
		respondWith(w, r, nil, ProductGroupsNotFoundMessage, nil, http.StatusNotFound, false)
		return
	}

	for _, result := range results {

		var pg PRODUCTGROUP
		mapDocument(w, r, &pg, result)
		productG = append(productG, pg)

	}

	respondWith(w, r, nil, ProductGroupsFoundMessage, productG, http.StatusOK, false)

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

	if !mapInput(w, r, &p) {
		return
	}

	if !validateProduct(w, r, p) {
		return
	}

	var opts options.FindOptions

	results := findMongoDocument(ExternalDB+csx, dbcol, bson.M{"Sku": p.Sku}, &opts)

	if len(results) != 0 {
		respondWith(w, r, nil, ProductAlreadyExistsMessage, nil, http.StatusConflict, false)
		return
	}

	groomProductData(&p)

	for _, cat := range p.Category {
		if !insertIntoTree(w, r, ExternalDB+csx, csx+CategoryTreeExtension, cat, p.Sku) {
			return
		}
	}

	p.Updated = time.Now().UnixNano()

	if !insertMongoDocument(ExternalDB+csx, dbcol, p) {
		respondWith(w, r, nil, HTTPInternalServerErrorMessage, nil, http.StatusInternalServerError, false)
		return
	}

	var productInventoryRecord INVENTORY
	productInventoryRecord.Sku = p.Sku
	productInventoryRecord.Quantity = p.Quantity
	productInventoryRecord.Updated = time.Now().UnixNano()

	if !insertMongoDocument(ExternalDB+csx, picol, productInventoryRecord) {
		respondWith(w, r, nil, HTTPInternalServerErrorMessage, nil, http.StatusInternalServerError, false)
		return
	}

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

	if !(mapInput(w, r, &p) && validateProduct(w, r, p)) {
		return
	}

	groomProductData(&p)

	csx := getAccessToken(r)
	dbcol := csx + ProductExtension
	ctcol := csx + CategoryTreeExtension

	var opts options.FindOptions

	results := findMongoDocument(ExternalDB+csx, dbcol, bson.M{"Sku": p.Sku}, &opts)

	if len(results) != 1 {
		respondWith(w, r, nil, ProductNotFoundMessage, nil, http.StatusNotFound, false)
		return
	}

	var rp PRODUCT

	mapDocument(w, r, &rp, results[0])

	for _, cat := range rp.Category {
		deleteSKUFromTree(w, r, ExternalDB+csx, ctcol, cat, rp.Sku)
	}

	result := updateMongoDocument(ExternalDB+csx, dbcol, bson.M{"Sku": p.Sku}, bson.M{"$set": p})

	if result[0] == 1 && result[1] == 1 {

		results := findMongoDocument(ExternalDB+csx, dbcol, bson.M{"Sku": p.Sku}, &opts)

		if len(results) != 1 {
			respondWith(w, r, nil, ProductNotFoundMessage, nil, http.StatusNotFound, false)
			return
		}

		var rp PRODUCT

		mapDocument(w, r, &rp, results[0])

		for _, cat := range p.Category {
			if !insertIntoTree(w, r, ExternalDB+csx, csx+CategoryTreeExtension, cat, p.Sku) {
				return
			}
		}

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

	results := findMongoDocument(ExternalDB+csx, dbcol, bson.M{"Sku": sku}, &opts)

	if len(results) != 1 {
		respondWith(w, r, nil, ProductNotFoundMessage, nil, http.StatusNotFound, false)
		return
	}

	var product PRODUCT

	mapDocument(w, r, &product, results[0])

	if deleteMongoDocument(ExternalDB+csx, dbcol, bson.M{"Sku": sku}) == 1 {

		deleteMongoDocument(ExternalDB+csx, picol, bson.M{"Sku": sku})

		for _, cat := range product.Category {
			deleteSKUFromTree(w, r, ExternalDB+csx, ctcol, cat, product.Sku)
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

		csx := getAccessToken(r)
		dbcol := csx + ProductGroupExtension

		pth := strings.Split(r.URL.Path, "/")
		pgid := pth[len(pth)-1]

		var opts options.FindOptions

		results := findMongoDocument(ExternalDB+csx, dbcol, bson.M{"GroupID": pgid}, &opts)

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

	results := findMongoDocument(ExternalDB+cidb, pgcol, bson.M{"GroupID": pgid}, &opts)

	if len(results) != 1 {
		respondWith(w, r, nil, ProductGroupNotFoundMessage, nil, http.StatusNotFound, false)
		return
	}

	var productGroup PRODUCTGROUP

	mapDocument(w, r, &productGroup, results[0])

	for _, product := range productGroup.Products {

		deleteMongoDocument(ExternalDB+cidb, pcol, bson.M{"Sku": product.Sku})
		deleteMongoDocument(ExternalDB+cidb, picol, bson.M{"Sku": product.Sku})

		for _, cat := range product.Category {
			deleteSKUFromTree(w, r, ExternalDB+cidb, ctcol, cat, product.Sku)
		}

	}

	if deleteMongoDocument(ExternalDB+cidb, pgcol, bson.M{"GroupID": pgid}) == 1 {

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

	if !mapInput(w, r, &prices) {
		return
	}

	for sku, price := range prices.Prices {
		if price.RegularPrice < 0 || price.PromotionPrice < 0 {
			respondWith(w, r, nil, "Price for sku: "+sku+" is negative. Prices cannot be negative ...", nil, http.StatusBadRequest, false)
			return
		}
	}

	csx := getAccessToken(r)
	dbcol := csx + ProductExtension

	var priceUpdated []string
	var priceNotUpdated []string
	var priceNotFound []string

	for sku, price := range prices.Prices {

		result := updateMongoDocument(ExternalDB+csx, dbcol, bson.M{"Sku": sku}, bson.M{"$set": bson.M{"RegularPrice": price.RegularPrice, "PromotionPrice": price.PromotionPrice}})

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

	if !mapInput(w, r, &quantities) {
		return
	}

	if len(quantities.Quantity) > InventoryUpdateBatchSize {
		respondWith(w, r, nil, "Inventory update batch size cannot be more than "+strconv.Itoa(InventoryUpdateBatchSize)+" per request ...", nil, http.StatusBadRequest, false)
		return
	}

	for sku, quantity := range quantities.Quantity {
		if quantity < 0 {
			respondWith(w, r, nil, "Inventory for sku: "+sku+" is negative. Quantity field cannot be negative ...", nil, http.StatusBadRequest, false)
			return
		}
	}

	csx := getAccessToken(r)
	picol := csx + ProductInventoryExtension

	var skusUpdated = make(map[string]int64)
	var skusNotUpdated []string
	var skusNotFound []string

	for sku, quantity := range quantities.Quantity {

		result := updateInventory(w, r, ExternalDB+csx, picol, "INCR", sku, quantity, true)

		if result[0] == 1 && result[1] == 1 {
			skusUpdated[sku] = quantities.Quantity[sku]
		} else if result[0] == 1 && result[1] == 0 {
			skusNotUpdated = append(skusNotUpdated, sku)
		} else if result[0] == 0 && result[1] == 0 {
			skusNotFound = append(skusNotFound, sku)
		}

	}

	respondWith(w, r, nil, "Inventory Update Status ...", bson.M{"Products Updated (Quantity mentioned below is incremented for each SKU)": skusUpdated, "Products Not Updated": skusNotUpdated, "Products Not Found": skusNotFound}, http.StatusOK, true)

}

func getProductInventory(w http.ResponseWriter, r *http.Request) {

	rlog.Debug("getProductInventory() handle function invoked ...")

	if !pre(w, r) {
		return
	}

	csx := getAccessToken(r)

	pth := strings.Split(r.URL.Path, "/")
	sku := pth[len(pth)-1]

	picol := csx + ProductInventoryExtension
	var opts options.FindOptions

	results := findMongoDocument(ExternalDB+csx, picol, bson.M{"Sku": sku}, &opts)

	if len(results) != 1 {
		respondWith(w, r, nil, "Inventory Record Not found ...", nil, http.StatusNotFound, false)
		return
	}

	var inventory INVENTORY

	mapDocument(w, r, &inventory, results[0])

	respondWith(w, r, nil, ProductFoundMessage, inventory, http.StatusOK, false)

}
