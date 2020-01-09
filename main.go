package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/magiconair/properties"
)

const properyFile = "/home/narang/work/src/github.com/viveknarang/kramaAPI/api.properties"

func main() {

	p := properties.MustLoadFile(properyFile, properties.UTF8)

	url := p.GetString("db.mongo.url", "localhost")
	port := p.GetString("db.mongo.port", "27017")
	apiPort := p.GetString("api.listen.on", "9005")
	catalogBasePath := p.GetString("api.catalog.base.path", "/catalog/")
	version := p.GetString("api.version", "v1")
	catalogPath := catalogBasePath + version

	connectDB(url, port)

	r := mux.NewRouter()

	r.HandleFunc("/customers/"+version+"/login", login).Methods(http.MethodPost)
	r.HandleFunc(catalogPath+"/products/{SKU}", getProduct).Methods(http.MethodGet)
	r.HandleFunc(catalogPath+"/products", postProduct).Methods(http.MethodPost)
	r.HandleFunc(catalogPath+"/products", putProduct).Methods(http.MethodPut)
	r.HandleFunc(catalogPath+"/products/{SKU}", deleteProduct).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe(":"+apiPort, r))

	disconnectDB()

}
