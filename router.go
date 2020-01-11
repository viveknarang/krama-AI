package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func routers() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/customers/"+APIVersion+"/login", login).Methods(http.MethodPost)

	router.HandleFunc(CatalogPath+"/products/{SKU}", getProduct).Methods(http.MethodGet)

	router.HandleFunc(CatalogPath+"/products", postProduct).Methods(http.MethodPost)

	router.HandleFunc(CatalogPath+"/products/{SKU}", putProduct).Methods(http.MethodPut)

	router.HandleFunc(CatalogPath+"/products/{SKU}", deleteProduct).Methods(http.MethodDelete)

	router.HandleFunc(CatalogPath+"/productgroups/{PGID}", getProductGroup).Methods(http.MethodGet)

	router.HandleFunc(CatalogPath+"/productgroups/{PGID}", deleteProductGroup).Methods(http.MethodDelete)

	return router

}
