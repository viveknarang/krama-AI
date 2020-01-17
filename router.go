package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/romana/rlog"
)

func routers() *mux.Router {

	rlog.Debug("Setting up routers and handle functions ...")

	router := mux.NewRouter()

	router.HandleFunc("/customers/"+APIVersion+"/login", login).Methods(http.MethodPost)

	router.HandleFunc(CatalogPath+"/products/{SKU}", getProduct).Methods(http.MethodGet)

	router.HandleFunc(CatalogPath+"/products", postProduct).Methods(http.MethodPost)

	router.HandleFunc(CatalogPath+"/products/{SKU}", putProduct).Methods(http.MethodPut)

	router.HandleFunc(CatalogPath+"/products/{SKU}", deleteProduct).Methods(http.MethodDelete)

	router.HandleFunc(CatalogPath+"/products/price/update", updateProductsPrice).Methods(http.MethodPut)

	router.HandleFunc(CatalogPath+"/products/inventory/update", updateProductsInventory).Methods(http.MethodPut)

	router.HandleFunc(CatalogPath+"/productgroups/{PGID}", getProductGroup).Methods(http.MethodGet)

	router.HandleFunc(CatalogPath+"/productgroups/{PGID}", deleteProductGroup).Methods(http.MethodDelete)

	router.HandleFunc(OrdersPath+"/orders/{OID}", getOrderByOrderID).Methods(http.MethodGet)

	router.HandleFunc(OrdersPath+"/orders/customer/{CID}", getOrderByCustomerID).Methods(http.MethodGet)

	router.HandleFunc(OrdersPath+"/orders", postOrder).Methods(http.MethodPost)

	router.HandleFunc(OrdersPath+"/orders/{OID}", putOrder).Methods(http.MethodPut)

	router.HandleFunc(OrdersPath+"/orders/{OID}", deleteOrder).Methods(http.MethodDelete)

	router.HandleFunc(SearchPath+"/quick", quickSearch).Methods(http.MethodPost)

	router.HandleFunc(SearchPath+"/fullpage", fullpageSearch).Methods(http.MethodPost)

	router.HandleFunc(CustomersPath+"/customers/{SKU}", getCustomer).Methods(http.MethodGet)

	router.HandleFunc(CustomersPath+"/customers", postCustomer).Methods(http.MethodPost)

	router.HandleFunc(CustomersPath+"/customers/{SKU}", putCustomer).Methods(http.MethodPut)

	router.HandleFunc(CustomersPath+"/customers/{SKU}", deleteCustomer).Methods(http.MethodDelete)

	rlog.Debug("Router setup complete ...")

	return router

}
