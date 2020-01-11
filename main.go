package main

import (
	"log"
	"net/http"
)

const properyFile = "/home/narang/work/src/github.com/viveknarang/kramaAPI/api.properties"

func main() {

	loadSystemProperties()

	connectDB(MongoURL, MongoPort)
	connectRedis(RedisURL, RedisPort)

	log.Fatal(http.ListenAndServe(":"+APIPort, routers()))

	disconnectDB()
	disconnectRedis()

}
