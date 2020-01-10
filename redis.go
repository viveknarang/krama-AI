package main

import (
	"github.com/go-redis/redis"
)

//REDISCLIENT client for redis
var REDISCLIENT *redis.Client

func connectRedis(url string, port string) {

	client := redis.NewClient(&redis.Options{
		Addr:     url + ":" + port,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()

	if pong == "PONG" && err == nil {
		REDISCLIENT = client
	} else {
		REDISCLIENT = nil
	}

}

func disconnectRedis() {
	REDISCLIENT.Close()
}
