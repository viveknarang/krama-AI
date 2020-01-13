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

	checkRedis := pingRedis()

	if checkRedis {
		fmt.Println("Redis connected at " + RedisURL + ":" + RedisPort)
	} else {
		return false
	}

	return checkRedis
}

func pingRedis() bool {

	var isRedisUp bool

	pong, err := REDISCLIENT.Ping().Result()

	if pong == "PONG" && err == nil {
		isRedisUp = true
	} else {
		isRedisUp = false
	}

	return isRedisUp

}

func disconnectRedis() {
	REDISCLIENT.Close()
}
