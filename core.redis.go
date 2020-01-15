package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

//REDISCLIENT client for redis
var REDISCLIENT *redis.Client

func connectRedis() bool {

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
	REDISCLIENT.Close()
}
