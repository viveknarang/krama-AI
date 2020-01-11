package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

//REDISCLIENT client for redis
var REDISCLIENT *redis.Client

func connectRedis(url string, port string) bool {

	var isRedisNormal bool

	client := redis.NewClient(&redis.Options{
		Addr:     url + ":" + port,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()

	if pong == "PONG" && err == nil {
		REDISCLIENT = client
		fmt.Println("Connected to Redis at " + url + ":" + port)
		isRedisNormal = true
	} else {
		REDISCLIENT = nil
		isRedisNormal = false
	}

	return isRedisNormal
}

func disconnectRedis() {
	REDISCLIENT.Close()
}
