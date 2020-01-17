package main

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/romana/rlog"
)

//REDISCLIENT client for redis
var REDISCLIENT *redis.Client

func connectRedis() bool {

	rlog.Debug("connectRedis() handle function invoked ...")

	client := redis.NewClient(&redis.Options{
		Addr:     RedisURL + ":" + RedisPort,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	REDISCLIENT = client

	checkRedis := pingRedis(false)

	return checkRedis
}

func pingRedis(silent bool) bool {

	rlog.Debug("pingRedis() handle function invoked ...")

	if REDISCLIENT == nil {
		return false
	}

	var isRedisUp bool

	pong, err := REDISCLIENT.Ping().Result()

	if pong == "PONG" && err == nil {
		isRedisUp = true
		if !silent {
			fmt.Println("ACTIVE PING FOR REDIS: Redis responding at " + RedisURL + ":" + RedisPort)
		}
	} else {
		isRedisUp = false
	}

	return isRedisUp

}

func disconnectRedis() {

	rlog.Debug("disconnectRedis() handle function invoked ...")

	REDISCLIENT.Close()
}
