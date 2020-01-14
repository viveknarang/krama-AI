package main

import (
	"log"
	"net/http"
)

func main() {

	loadSystemProperties()

	connectDB()
	connectRedis()
	connectElastic()

	log.Fatal(http.ListenAndServe(":"+APIPort, routers()))

	disconnectDB()
	disconnectRedis()

	some garbage here to test travis-ci
	
}
