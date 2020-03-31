package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"github.com/romana/rlog"
)

func throttleCheck(w http.ResponseWriter, r *http.Request, authToken string) bool {

	if REDISCLIENT.Get(authToken+"_rxt").Err() != redis.Nil {

		throttleRate, _ := strconv.Atoi(REDISCLIENT.Get(authToken + "_rxt").Val())

		currentMinute := time.Now().Unix() / 60

		rateKey := authToken + ":" + strconv.FormatInt(currentMinute, 10)

		key := REDISCLIENT.Get(rateKey)

		if key.Err() != redis.Nil {

			count, _ := strconv.Atoi(key.Val())

			if count >= throttleRate {

				//Return Error
				rlog.Debug("pre()/throttle check: Request rate capacity is reached for this customer ...")
				respondWith(w, r, nil, "Request rate capacity of "+strconv.Itoa(throttleRate)+" requests per minute is reached for this customer ...", nil, http.StatusTooManyRequests, false)
				return false

			}

		}

		REDISCLIENT.Incr(rateKey)
		REDISCLIENT.Expire(rateKey, 5.9e+10)

	}

	return true

}
