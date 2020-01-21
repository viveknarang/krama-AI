package main

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/romana/rlog"
	"go.mongodb.org/mongo-driver/bson"
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

func fileExists(fileName string) bool {

	if _, err := os.Stat(fileName); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true

}

func typeof(value interface{}) string {

	return reflect.TypeOf(value).String()

}

func areCoreServicesUp() bool {

	rlog.Debug("areCoreServicesUp() function invoked ...")

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

func mapInput(w http.ResponseWriter, r *http.Request, object interface{}) {

	err := json.NewDecoder(r.Body).Decode(&object)

	if err != nil {

		respondWith(w, r, err, HTTPBadRequestMessage, nil, http.StatusBadRequest, false)
		return

	}

}

func mapDocument(w http.ResponseWriter, r *http.Request, object interface{}, document interface{}) {

	j, err0 := bson.MarshalExtJSON(document, false, false)

	if err0 != nil {
		respondWith(w, r, err0, HTTPInternalServerErrorMessage, nil, http.StatusInternalServerError, false)
		return
	}

	err1 := json.Unmarshal([]byte(j), &object)

	if err1 != nil {
		respondWith(w, r, err1, HTTPInternalServerErrorMessage, nil, http.StatusInternalServerError, false)
		return
	}

}

func mapBytes(w http.ResponseWriter, r *http.Request, object interface{}, bytes []byte) {

	err1 := json.Unmarshal(bytes, &object)

	if err1 != nil {
		respondWith(w, r, err1, HTTPInternalServerErrorMessage, nil, http.StatusInternalServerError, false)
		return
	}

}

func mapToBytes(w http.ResponseWriter, r *http.Request, document interface{}) []byte {

	j, err0 := bson.MarshalExtJSON(document, false, false)

	if err0 != nil {
		respondWith(w, r, err0, HTTPInternalServerErrorMessage, nil, http.StatusInternalServerError, false)
		return nil
	}

	return j

}

func getAccessToken(r *http.Request) string {

	return REDISCLIENT.Get(r.Header.Get("x-access-token")).Val()

}

func logDebugMessage(message string) {

	rlog.Debug(message)

}

func logErrorMessage(message string) {

	rlog.Error(message)

}

func logInfoMessage(message string) {

	rlog.Info(message)

}

func customValidatorForTypeArrayLengths(v interface{}, param string) error {

	st := reflect.ValueOf(v)

	iparam, _ := strconv.Atoi(param)

	if st.Len() > int(iparam) {
		return errors.New("field has more than " + param + " entries")
	}

	return nil

}

func customValidatorForEmail(v interface{}, param string) error {

	st := reflect.ValueOf(v)

	if !isValidEmail(st.String()) {
		return errors.New("field is not a valid email address")
	}

	return nil
}

func customValidatorForNoSpaces(v interface{}, param string) error {

	st := reflect.ValueOf(v)

	if strings.ContainsAny(st.String(), " ") {
		return errors.New("field should not have spaces it it")
	}

	return nil
}

func customValidatorForMaxFloat(v interface{}, param string) error {

	st := reflect.ValueOf(v)

	if st.Float() > math.MaxFloat64 {
		return errors.New("field has a value bigger than what the system can handle")
	}

	return nil
}

func customValidatorForAllowedCurrencies(v interface{}, param string) error {

	st := reflect.ValueOf(v)

	var currency [6]string

	//This list will add more values in the near future
	currency[0] = "USD"
	currency[1] = "CAD"
	currency[2] = "CDN"
	currency[3] = "EUR"
	currency[4] = "INR"
	currency[5] = "GBP"

	if !contains(currency, st.String()) {
		return errors.New("field value " + st.String() + " not an acceptable currency")
	}

	return nil
}

func contains(arr [6]string, str string) bool {

	for _, a := range arr {
		if a == str {
			return true
		}
	}

	return false

}
