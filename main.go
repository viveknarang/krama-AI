package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/romana/rlog"
)

func main() {

	loadEnvironmentVariables()

	if !loadSystemProperties() {
		return
	}

	rlog.Debug("Attempting to connect to base components - [ELASTIC, REDIS, MONGO]...")

	connectDB()
	connectRedis()
	connectElastic()
	setCustomValidators()

	rlog.Debug("Base components - [ELASTIC, REDIS, MONGO] connected ...")

	headers := handlers.AllowedHeaders([]string{"x-access-token", "X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	log.Fatal(http.ListenAndServe(":"+APIPort, handlers.CORS(headers, methods, origins)(routers())))

}
