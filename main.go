package main

import (
	"log"
	"net/http"
)

func main() {

	loadSystemProperties()

	connectDB(MongoURL, MongoPort)
	connectRedis(RedisURL, RedisPort)

	log.Fatal(http.ListenAndServe(":"+APIPort, routers()))

	disconnectDB()
	disconnectRedis()

}
