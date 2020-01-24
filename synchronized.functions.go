package main

import (
	"net/http"
	"sync"
	"time"

	"github.com/romana/rlog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mutex sync.Mutex

// Synchronized function to keep inventory levels consistent...
func updateInventory(w http.ResponseWriter, r *http.Request, collection string, iodi string, Sku string, count int64, ignoreMessageForNotFound bool) [2]int64 {

	rlog.Debug("updateInventory() handle function invoked ...")

	mutex.Lock()

	var fr [2]int64
	fr[0] = -1
	fr[1] = -1

	var opts options.FindOptions

	results := findMongoDocument(ExternalDB, collection, bson.M{"Sku": Sku}, &opts)

	if len(results) != 1 {

		if !ignoreMessageForNotFound {
			respondWith(w, r, nil, "Inventory Record Not found ...", nil, http.StatusNotFound, false)
		}

		mutex.Unlock()

		fr[0] = 0
		fr[1] = 0

		return fr
	}

	var productInventoryRecord INVENTORY

	mapDocument(w, r, &productInventoryRecord, results[0])

	if iodi == "DECR" {

		if productInventoryRecord.Quantity <= 0 || productInventoryRecord.Quantity-count <= 0 {
			respondWith(w, r, nil, "Product with SKU: "+productInventoryRecord.Sku+" is either out of stock or not enough stock for fulfilling your order ...", nil, http.StatusNotFound, false)
			mutex.Unlock()
			return fr
		}

		productInventoryRecord.Quantity = productInventoryRecord.Quantity - count

	} else if iodi == "INCR" {

		productInventoryRecord.Quantity = productInventoryRecord.Quantity + count

	}

	productInventoryRecord.Updated = time.Now().UnixNano()

	result := updateMongoDocument(ExternalDB, collection, bson.M{"Sku": productInventoryRecord.Sku}, bson.M{"$set": productInventoryRecord})

	mutex.Unlock()

	return result

}
