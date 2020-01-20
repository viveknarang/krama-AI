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

func getShoppingCart(w http.ResponseWriter, r *http.Request) {

	rlog.Debug("getShoppingCart() handle function invoked ...")

	if !pre(w, r) {
		return
	}

	pth := strings.Split(r.URL.Path, "/")
	cid := pth[len(pth)-1]

	shoppingCartO := REDISCLIENT.Get(cid)

	if shoppingCartO.Val() == "" {
		respondWith(w, r, nil, "Cart id: "+cid+" not found ...", nil, http.StatusNotFound, false)
		return
	}

	jx := []byte(shoppingCartO.Val())

	var shoppingCart SHOPPINGCART

	err1 := json.Unmarshal([]byte(jx), &shoppingCart)

	if err1 != nil {
		respondWith(w, r, err1, HTTPInternalServerErrorMessage, nil, http.StatusInternalServerError, false)
		return
	}

	respondWith(w, r, nil, "Shopping Cart: ", shoppingCart, http.StatusOK, true)

}

func addProductInShoppingCart(w http.ResponseWriter, r *http.Request) {

	rlog.Debug("addProductInShoppingCart() handle function invoked ...")

	if !pre(w, r) {
		return
	}

	var shoppingCartReq SHOPPINGCARTREQ

	err := json.NewDecoder(r.Body).Decode(&shoppingCartReq)

	if err != nil {
		respondWith(w, r, err, HTTPBadRequestMessage, nil, http.StatusBadRequest, false)
		return
	}

	shoppingCartO := REDISCLIENT.Get(shoppingCartReq.CartID)
	var shoppingCart SHOPPINGCART

	if shoppingCartO.Val() != "" {

		jx := []byte(shoppingCartO.Val())

		err1 := json.Unmarshal([]byte(jx), &shoppingCart)

		if err1 != nil {
			respondWith(w, r, err1, HTTPInternalServerErrorMessage, nil, http.StatusInternalServerError, false)
			return
		}

	}

	csx := REDISCLIENT.Get(r.Header.Get("x-access-token")).Val()
	picol := csx + ProductInventoryExtension
	var opts options.FindOptions

	results := findMongoDocument(ExternalDB, picol, bson.M{"Sku": shoppingCartReq.Product.Sku}, &opts)

	if len(results) != 1 {
		respondWith(w, r, nil, "Inventory Record Not found ...", nil, http.StatusNotFound, false)
		return
	}

	j, err0 := bson.MarshalExtJSON(results[0], false, false)

	if err0 != nil {
		respondWith(w, r, err0, HTTPInternalServerErrorMessage, nil, http.StatusInternalServerError, false)
		return
	}

	var productInventoryRecord INVENTORY

	err3 := json.Unmarshal([]byte(j), &productInventoryRecord)

	if err3 != nil {
		respondWith(w, r, err3, HTTPInternalServerErrorMessage, nil, http.StatusInternalServerError, false)
		return
	}

	if productInventoryRecord.Quantity <= 0 || productInventoryRecord.Quantity-shoppingCartReq.Count <= 0 {
		respondWith(w, r, nil, "Out of stock ...", nil, http.StatusNotFound, false)
		return
	}

	productInventoryRecord.Quantity = productInventoryRecord.Quantity - shoppingCartReq.Count
	productInventoryRecord.Updated = time.Now().UnixNano()

	result := updateMongoDocument(ExternalDB, picol, bson.M{"Sku": productInventoryRecord.Sku}, bson.M{"$set": productInventoryRecord})

	if result[1] == 0 {
		respondWith(w, r, nil, HTTPInternalServerErrorMessage, nil, http.StatusInternalServerError, false)
		return
	}

	if shoppingCartReq.CartID == "" {
		shoppingCart.CartID = uuid.New().String()
	} else {
		shoppingCart.CartID = shoppingCartReq.CartID
	}

	if shoppingCart.Products == nil {
		shoppingCart.Products = make(map[string]PRODUCT)
	}

	if shoppingCart.ProductsCount == nil {
		shoppingCart.ProductsCount = make(map[string]int64)
	}

	if shoppingCartReq.CustomerID != "" {
		shoppingCart.CustomerID = shoppingCartReq.CustomerID
	}

	shoppingCart.ProductsCount[shoppingCartReq.Product.Sku] += shoppingCartReq.Count
	shoppingCart.Products[shoppingCartReq.Product.Sku] = shoppingCartReq.Product
	shoppingCart.Updated = time.Now().UnixNano()
	shoppingCart.Currency = shoppingCartReq.Product.Currency

	var total float64
	for key, value := range shoppingCart.Products {

		if value.PromotionPrice < value.RegularPrice {
			total += value.PromotionPrice * float64(shoppingCart.ProductsCount[key])
		} else {
			total += value.RegularPrice * float64(shoppingCart.ProductsCount[key])
		}

	}

	shoppingCart.Total = total

	bt, _ := json.Marshal(shoppingCart)

	REDISCLIENT.Set(shoppingCart.CartID, bt, 0)

	respondWith(w, r, nil, "Product added in the cart...", shoppingCart, http.StatusOK, true)

}

func removeProductFromShoppingCart(w http.ResponseWriter, r *http.Request) {

	rlog.Debug("removeProductFromShoppingCart() handle function invoked ...")

	if !pre(w, r) {
		return
	}

	var shoppingCartReq SHOPPINGCARTREQ

	err := json.NewDecoder(r.Body).Decode(&shoppingCartReq)

	if err != nil {
		respondWith(w, r, err, HTTPBadRequestMessage, nil, http.StatusBadRequest, false)
		return
	}

	shoppingCartO := REDISCLIENT.Get(shoppingCartReq.CartID)
	var shoppingCart SHOPPINGCART

	if shoppingCartO.Val() != "" {

		jx := []byte(shoppingCartO.Val())

		err1 := json.Unmarshal([]byte(jx), &shoppingCart)

		if err1 != nil {
			respondWith(w, r, err1, HTTPInternalServerErrorMessage, nil, http.StatusInternalServerError, false)
			return
		}

	} else {

		respondWith(w, r, nil, "Cart id: "+shoppingCartReq.CartID+" not found ...", nil, http.StatusNotFound, false)
		return

	}

	if shoppingCart.Products == nil {
		shoppingCart.Products = make(map[string]PRODUCT)
	}

	if shoppingCart.ProductsCount == nil {
		shoppingCart.ProductsCount = make(map[string]int64)
	}

	if shoppingCartReq.CustomerID != "" {
		shoppingCart.CustomerID = shoppingCartReq.CustomerID
	}

	csx := REDISCLIENT.Get(r.Header.Get("x-access-token")).Val()
	picol := csx + ProductInventoryExtension
	var opts options.FindOptions

	results := findMongoDocument(ExternalDB, picol, bson.M{"Sku": shoppingCartReq.SKU}, &opts)

	if len(results) != 1 {
		respondWith(w, r, nil, "Inventory Record Not found ...", nil, http.StatusNotFound, false)
		return
	}

	j, err0 := bson.MarshalExtJSON(results[0], false, false)

	if err0 != nil {
		respondWith(w, r, err0, HTTPInternalServerErrorMessage, nil, http.StatusInternalServerError, false)
		return
	}

	var productInventoryRecord INVENTORY

	err3 := json.Unmarshal([]byte(j), &productInventoryRecord)

	if err3 != nil {
		respondWith(w, r, err3, HTTPInternalServerErrorMessage, nil, http.StatusInternalServerError, false)
		return
	}

	productInventoryRecord.Quantity = productInventoryRecord.Quantity + shoppingCartReq.Count
	productInventoryRecord.Updated = time.Now().UnixNano()

	result := updateMongoDocument(ExternalDB, picol, bson.M{"Sku": productInventoryRecord.Sku}, bson.M{"$set": productInventoryRecord})

	if result[1] == 0 {
		respondWith(w, r, nil, HTTPInternalServerErrorMessage, nil, http.StatusInternalServerError, false)
		return
	}

	shoppingCart.ProductsCount[shoppingCartReq.SKU] -= shoppingCartReq.Count

	if shoppingCart.ProductsCount[shoppingCartReq.SKU] <= 0 {
		delete(shoppingCart.ProductsCount, shoppingCartReq.SKU)
		delete(shoppingCart.Products, shoppingCartReq.SKU)
	}

	shoppingCart.Updated = time.Now().UnixNano()

	var total float64
	for key, value := range shoppingCart.Products {

		if value.PromotionPrice < value.RegularPrice {
			total += value.PromotionPrice * float64(shoppingCart.ProductsCount[key])
		} else {
			total += value.RegularPrice * float64(shoppingCart.ProductsCount[key])
		}

	}

	shoppingCart.Total = total

	bt, _ := json.Marshal(shoppingCart)

	REDISCLIENT.Set(shoppingCart.CartID, bt, 0)

	respondWith(w, r, nil, "Product removed the cart...", shoppingCart, http.StatusOK, true)
}

func clearShoppingCart(w http.ResponseWriter, r *http.Request) {

	rlog.Debug("clearShoppingCart() handle function invoked ...")

	if !pre(w, r) {
		return
	}

	pth := strings.Split(r.URL.Path, "/")
	cid := pth[len(pth)-1]

	shoppingCartO := REDISCLIENT.Get(cid)

	if shoppingCartO.Val() != "" {

		REDISCLIENT.Del(cid)
		respondWith(w, r, nil, "Cart with id: "+cid+" deleted ...", nil, http.StatusAccepted, true)

	} else {

		respondWith(w, r, nil, "Cart id: "+cid+" not found ...", nil, http.StatusNotFound, false)

	}

}
