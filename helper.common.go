package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"regexp"
)

func isValidJSON(s string) bool {

	var js interface{}
	return json.Unmarshal([]byte(s), &js) == nil

}

func isValidURL(toTest string) bool {

	_, err := url.ParseRequestURI(toTest)

	return !(err != nil)

}

func isValidEmail(toTest string) bool {

	emailRegularExpression := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	return emailRegularExpression.MatchString(toTest)

}

func isValidAttributeKey(potentialAttributeKey string) bool {

	attributeKey := regexp.MustCompile("^[A-Za-z0-9]+([-_ ]{1}[A-Za-z0-9]+)*$")

	return attributeKey.MatchString(potentialAttributeKey)

}

func hashString(Txt string) string {
	h := sha256.New()
	h.Write([]byte(Txt))
	bs := h.Sum(nil)
	sh := string(fmt.Sprintf("%x", bs))
	return sh
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
