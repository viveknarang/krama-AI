package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const properyFile = "/home/narang/work/src/github.com/viveknarang/kramaAPI/api.properties"

func main() {

	loadSystemProperties()

	connectDB(MongoURL, MongoPort)
	connectRedis(RedisURL, RedisPort)

	r := mux.NewRouter()

	r.HandleFunc("/customers/"+APIVersion+"/login", login).Methods(http.MethodPost)
	r.HandleFunc(CatalogPath+"/products/{SKU}", getProduct).Methods(http.MethodGet)
	r.HandleFunc(CatalogPath+"/products", postProduct).Methods(http.MethodPost)
	r.HandleFunc(CatalogPath+"/products", putProduct).Methods(http.MethodPut)
	r.HandleFunc(CatalogPath+"/products/{SKU}", deleteProduct).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe(":"+APIPort, r))

	disconnectDB()
	disconnectRedis()

}
