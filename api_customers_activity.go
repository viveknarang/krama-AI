package main

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/go-redis/redis"
	"github.com/romana/rlog"
)

func getCustomerBrowsingHistory(w http.ResponseWriter, r *http.Request) {

	rlog.Debug("getCustomerBrowsingHistory() handle function invoked ...")

	if !pre(w, r) {
		return
	}

	pth := strings.Split(r.URL.Path, "/")
	length := len(pth)

	customerID := pth[len(pth)-3]

	if len(customerID) > 50 || len(pth[length-2]) > 50 || len(pth[length-1]) > 50 {
		respondWith(w, r, nil, "Bad Request", "Bad Request! Possibly too long inputs. Please don't mess with the platform!", http.StatusBadRequest, false)
		return
	}

	start, err0 := strconv.ParseInt(pth[length-2], 10, 64)

	limit, err1 := strconv.ParseInt(pth[length-1], 10, 64)

	browsingHistory := REDISCLIENT.ZRange(customerID, start, limit)

	if err0 == nil && err1 == nil {
		respondWith(w, r, nil, "Browsing History", bson.M{customerID: browsingHistory.Val()}, http.StatusOK, true)
	} else {
		respondWith(w, r, nil, "Bad Request", "The request is malformed! Please make sure that the range attributes in the request URL are integers", http.StatusBadRequest, false)
	}

}

func postCustomerBrowsingHistory(w http.ResponseWriter, r *http.Request) {

	rlog.Debug("postCustomerBrowsingHistory() handle function invoked ...")

	if !pre(w, r) {
		return
	}

	pth := strings.Split(r.URL.Path, "/")

	pgid := pth[len(pth)-1]

	customerID := pth[len(pth)-2]

	time := time.Now().Unix()

	REDISCLIENT.ZRem(customerID, pgid)

	response := REDISCLIENT.ZAdd(customerID, &redis.Z{
		Score:  float64(time),
		Member: pgid,
	})

	if response.Val() == 1 {
		respondWith(w, r, nil, "Browsing History Added", bson.M{"Product": pgid, "Customer": customerID, "Time": time}, http.StatusOK, true)
	} else {
		respondWith(w, r, nil, "Internal Server Error", "Something went wrong!", http.StatusInternalServerError, false)
	}
}
