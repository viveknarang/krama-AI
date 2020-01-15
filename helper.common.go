package main

import (
	"encoding/json"
	"net/url"
	"reflect"
)

func isValidJSON(s string) bool {

	var js interface{}
	return json.Unmarshal([]byte(s), &js) == nil

}

func isValidURL(toTest string) bool {

	_, err := url.ParseRequestURI(toTest)

	return !(err != nil)

}

func typeof(value interface{}) string {

	return reflect.TypeOf(value).String()

}

func areCoreServicesUp() bool {

	return pingMongoDB(true) && pingES(true) && pingRedis(true)

}

func resetProductCacheKeys(p *PRODUCT, pg *PRODUCTGROUP) {

	if p != nil {
		REDISCLIENT.Del(CatalogPath + "/products/" + p.Sku)
		REDISCLIENT.Del(CatalogPath + "/productgroups/" + p.GroupID)
	} else if pg != nil {
		REDISCLIENT.Del(CatalogPath + "/productgroups/" + pg.GroupID)
		for _, sku := range pg.Skus {
			REDISCLIENT.Del(CatalogPath + "/products/" + sku)
		}
	}

}

func resetCustomerCacheKeys(customer *CUSTOMER) {

	if customer != nil {
		REDISCLIENT.Del(CustomersPath + "/customers/" + customer.CustomerID)
	}

}
