package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"github.com/romana/rlog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getCustomer(w http.ResponseWriter, r *http.Request) {

	rlog.Debug("getCustomer() handle function invoked ...")

	if !pre(w, r) {
		return
	}

	var jx []byte
	var customer CUSTOMER

	redisC := REDISCLIENT.Get(r.URL.Path)

	if redisC.Err() != redis.Nil {

		jx = []byte(redisC.Val())
		mapBytes(w, r, &customer, jx)

	} else {

		pth := strings.Split(r.URL.Path, "/")
		cid := pth[len(pth)-1]

		dbcol := getAccessToken(r) + CustomersCollectionExtension

		var opts options.FindOptions

		results := findMongoDocument(ExternalDB, dbcol, bson.M{"CustomerID": cid}, &opts)

		if len(results) != 1 {
			respondWith(w, r, nil, CustomersNotFoundMessage, nil, http.StatusNotFound, false)
			return
		}

		mapDocument(w, r, &customer, results[0])

		jx = mapToBytes(w, r, results[0])

		REDISCLIENT.Set(r.URL.Path, jx, 0)

	}

	respondWith(w, r, nil, CustomersFoundMessage, customer, http.StatusOK, false)

}

func postCustomer(w http.ResponseWriter, r *http.Request) {

	rlog.Debug("postCustomer() handle function invoked ...")

	if !pre(w, r) {
		return
	}

	var customer CUSTOMER

	if !mapInput(w, r, &customer) {
		return
	}

	dbcol := getAccessToken(r) + CustomersCollectionExtension

	var opts options.FindOptions

	results := findMongoDocument(ExternalDB, dbcol, bson.M{"Email": customer.Email}, &opts)

	if len(results) != 0 {
		respondWith(w, r, nil, CustomerAlreadyExistsMessage, nil, http.StatusConflict, false)
		return
	}

	if customer.CustomerID == "" {
		customer.CustomerID = uuid.New().String()
	}

	if !validateCustomer(w, r, customer) {
		return
	}

	groomCustomerData(&customer)

	customer.Password = hashString(customer.Password)
	customer.Updated = time.Now().UnixNano()

	insertMongoDocument(ExternalDB, dbcol, customer)

	respondWith(w, r, nil, CustomersAddedMessage, customer, http.StatusCreated, true)

}

func putCustomer(w http.ResponseWriter, r *http.Request) {

	rlog.Debug("putCustomer() handle function invoked ...")

	if !pre(w, r) {
		return
	}

	pth := strings.Split(r.URL.Path, "/")
	cid := pth[len(pth)-1]

	var customer CUSTOMER

	if !mapInput(w, r, &customer) {
		return
	}

	customer.CustomerID = cid

	if !validateCustomer(w, r, customer) {
		return
	}

	groomCustomerData(&customer)

	customer.Updated = time.Now().UnixNano()

	dbcol := getAccessToken(r) + CustomersCollectionExtension

	result := updateMongoDocument(ExternalDB, dbcol, bson.M{"CustomerID": customer.CustomerID}, bson.M{"$set": customer})

	if result[0] == 1 && result[1] == 1 {

		resetCustomerCacheKeys(&customer)
		respondWith(w, r, nil, CustomersUpdatedMessage, customer, http.StatusAccepted, true)

	} else if result[0] == 1 && result[1] == 0 {

		respondWith(w, r, nil, CustomersNotUpdatedMessage, nil, http.StatusNotModified, false)

	} else if result[0] == 0 && result[1] == 0 {

		respondWith(w, r, nil, CustomersNotFoundMessage, nil, http.StatusNotModified, false)

	}

}

func deleteCustomer(w http.ResponseWriter, r *http.Request) {

	rlog.Debug("deleteCustomer() handle function invoked ...")

	if !pre(w, r) {
		return
	}

	dbcol := getAccessToken(r) + CustomersCollectionExtension

	pth := strings.Split(r.URL.Path, "/")
	cid := pth[len(pth)-1]

	var opts options.FindOptions

	results := findMongoDocument(ExternalDB, dbcol, bson.M{"CustomerID": cid}, &opts)

	if len(results) != 1 {
		respondWith(w, r, nil, CustomersNotFoundMessage, nil, http.StatusNotFound, false)
		return
	}

	var customer CUSTOMER

	mapDocument(w, r, &customer, results[0])

	if deleteMongoDocument(ExternalDB, dbcol, bson.M{"CustomerID": cid}) == 1 {

		resetCustomerCacheKeys(&customer)
		respondWith(w, r, nil, CustomersDeletedMessage, nil, http.StatusOK, true)

	} else {

		respondWith(w, r, nil, CustomersNotFoundMessage, nil, http.StatusNotModified, false)

	}

}
