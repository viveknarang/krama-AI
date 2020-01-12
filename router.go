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

	router.HandleFunc(OrdersPath+"/orders/{OID}", getOrderByOrderID).Methods(http.MethodGet)

	router.HandleFunc(OrdersPath+"/orders/customer/{CID}", getOrderByCustomerID).Methods(http.MethodGet)

	router.HandleFunc(OrdersPath+"/orders", postOrder).Methods(http.MethodPost)

	router.HandleFunc(OrdersPath+"/orders/{SKU}", putOrder).Methods(http.MethodPut)

	router.HandleFunc(OrdersPath+"/orders/{SKU}", deleteOrder).Methods(http.MethodDelete)

	router.HandleFunc(SearchPath+"/productgroups/search", basicProductGroupSearch).Methods(http.MethodPost)

	return router

}
