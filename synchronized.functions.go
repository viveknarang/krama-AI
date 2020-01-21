package main

import (
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Synchronized function to keep inventory levels consistent...
func updateInventory(w http.ResponseWriter, r *http.Request, collection string, iodi string, Sku string, count int64) {

	mutex.Lock()

	var opts options.FindOptions

	results := findMongoDocument(ExternalDB, collection, bson.M{"Sku": Sku}, &opts)

	if len(results) != 1 {
		respondWith(w, r, nil, "Inventory Record Not found ...", nil, http.StatusNotFound, false)
		mutex.Unlock()
		return
	}

	var productInventoryRecord INVENTORY

	mapDocument(w, r, &productInventoryRecord, results[0])

	if iodi == "DECR" {

		if productInventoryRecord.Quantity <= 0 || productInventoryRecord.Quantity-count <= 0 {
			respondWith(w, r, nil, "Product with SKU: "+productInventoryRecord.Sku+" is either out of stock or not enough stock for fulfilling your order ...", nil, http.StatusNotFound, false)
			mutex.Unlock()
			return
		}

		productInventoryRecord.Quantity = productInventoryRecord.Quantity - count

	} else if iodi == "INCR" {

		productInventoryRecord.Quantity = productInventoryRecord.Quantity + count

	}

	productInventoryRecord.Updated = time.Now().UnixNano()

	result := updateMongoDocument(ExternalDB, collection, bson.M{"Sku": productInventoryRecord.Sku}, bson.M{"$set": productInventoryRecord})

	if result[1] == 0 {
		respondWith(w, r, nil, HTTPInternalServerErrorMessage, nil, http.StatusInternalServerError, false)
		mutex.Unlock()
		return
	}

	mutex.Unlock()

}
