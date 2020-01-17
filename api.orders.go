package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"github.com/romana/rlog"
	"go.mongodb.org/mongo-driver/bson"
)

func getOrderByOrderID(w http.ResponseWriter, r *http.Request) {

	rlog.Debug("getOrderByOrderID() handle function invoked ...")

	if !pre(w, r) {
		return
	}

	var jx []byte

	redisC := REDISCLIENT.Get(r.URL.Path)

	if redisC.Err() != redis.Nil {

		jx = []byte(redisC.Val())

	} else {

		pth := strings.Split(r.URL.Path, "/")
		oid := pth[len(pth)-1]

		dbcol := REDISCLIENT.Get(r.Header.Get("x-access-token")).Val() + OrdersExtension

		results := findMongoDocument(ExternalDB, dbcol, bson.M{"orderid": oid})

		if len(results) != 1 {
			respondWith(w, r, nil, OrderNotFoundMessage, nil, http.StatusNotFound, false)
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

	var order ORDER

	err1 := json.Unmarshal([]byte(jx), &order)

	if err1 != nil {
		respondWith(w, r, err1, HTTPInternalServerErrorMessage, nil, http.StatusInternalServerError, false)
		return
	}

	respondWith(w, r, nil, OrderFoundMessage, order, http.StatusOK, true)

}

func getOrderByCustomerID(w http.ResponseWriter, r *http.Request) {

	rlog.Debug("getOrderByCustomerID() handle function invoked ...")

	if !pre(w, r) {
		return
	}

	var jx []byte

	redisC := REDISCLIENT.Get(r.URL.Path)

	if redisC.Err() != redis.Nil {

		jx = []byte(redisC.Val())

	} else {

		pth := strings.Split(r.URL.Path, "/")
		cid := pth[len(pth)-1]

		dbcol := REDISCLIENT.Get(r.Header.Get("x-access-token")).Val() + OrdersExtension

		results := findMongoDocument(ExternalDB, dbcol, bson.M{"customerid": cid})

		if len(results) != 1 {
			respondWith(w, r, nil, OrderNotFoundMessage, nil, http.StatusNotFound, false)
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

	var order ORDER

	err1 := json.Unmarshal([]byte(jx), &order)

	if err1 != nil {
		respondWith(w, r, err1, HTTPInternalServerErrorMessage, nil, http.StatusInternalServerError, false)
		return
	}

	respondWith(w, r, nil, OrderFoundMessage, order, http.StatusOK, true)

}

func postOrder(w http.ResponseWriter, r *http.Request) {

	rlog.Debug("postOrder() handle function invoked ...")

	if !pre(w, r) {
		return
	}

	var order ORDER

	err := json.NewDecoder(r.Body).Decode(&order)

	if err != nil {
		respondWith(w, r, err, HTTPBadRequestMessage, nil, http.StatusBadRequest, false)
		return
	}

	order.OrderCreationDate = time.Now().UnixNano()

	if order.OrderID == "" {
		order.OrderID = uuid.New().String()
	}

	dbcol := REDISCLIENT.Get(r.Header.Get("x-access-token")).Val() + OrdersExtension

	insertMongoDocument(ExternalDB, dbcol, order)

	respondWith(w, r, nil, OrderCreatedMessage, order, http.StatusCreated, true)

}

func putOrder(w http.ResponseWriter, r *http.Request) {

	rlog.Debug("putOrder() handle function invoked ...")

	if !pre(w, r) {
		return
	}

	pth := strings.Split(r.URL.Path, "/")
	oid := pth[len(pth)-1]

	var order ORDER

	err := json.NewDecoder(r.Body).Decode(&order)

	if err != nil {
		respondWith(w, r, err, HTTPBadRequestMessage, nil, http.StatusBadRequest, false)
		return
	}

	order.OrderCreationDate = time.Now().UnixNano()
	order.OrderID = oid

	dbcol := REDISCLIENT.Get(r.Header.Get("x-access-token")).Val() + OrdersExtension

	result := updateMongoDocument(ExternalDB, dbcol, bson.M{"orderid": order.OrderID}, bson.M{"$set": order})

	if result[0] == 1 && result[1] == 1 {

		REDISCLIENT.Del(r.URL.Path)
		respondWith(w, r, nil, OrderUpdatedMessage, order, http.StatusAccepted, true)

	} else if result[0] == 1 && result[1] == 0 {

		respondWith(w, r, nil, OrderNotUpdatedMessage, nil, http.StatusNotModified, false)

	} else if result[0] == 0 && result[1] == 0 {

		respondWith(w, r, nil, OrderNotFoundMessage, nil, http.StatusNotModified, false)

	}

}

func deleteOrder(w http.ResponseWriter, r *http.Request) {

	rlog.Debug("deleteOrder() handle function invoked ...")

	if !pre(w, r) {
		return
	}

	dbcol := REDISCLIENT.Get(r.Header.Get("x-access-token")).Val() + OrdersExtension

	pth := strings.Split(r.URL.Path, "/")
	oid := pth[len(pth)-1]

	results := findMongoDocument(ExternalDB, dbcol, bson.M{"orderid": oid})

	if len(results) != 1 {
		respondWith(w, r, nil, OrderNotFoundMessage, nil, http.StatusNotFound, false)
		return
	}

	j, err0 := bson.MarshalExtJSON(results[0], false, false)

	if err0 != nil {
		respondWith(w, r, err0, HTTPInternalServerErrorMessage, nil, http.StatusInternalServerError, false)
		return
	}

	var order ORDER

	err1 := json.Unmarshal([]byte(j), &order)

	if err1 != nil {
		respondWith(w, r, err1, HTTPInternalServerErrorMessage, nil, http.StatusInternalServerError, false)
		return
	}

	if deleteMongoDocument(ExternalDB, dbcol, bson.M{"orderid": oid}) == 1 {

		REDISCLIENT.Del(r.URL.Path)
		respondWith(w, r, nil, OrderDeletedMessage, nil, http.StatusOK, true)

	} else {

		respondWith(w, r, nil, OrderNotFoundMessage, nil, http.StatusNotModified, false)

	}

}
