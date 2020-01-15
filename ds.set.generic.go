package main

import (
	"reflect"
)

var gmp = make(map[interface{}]bool)

func addInGSet(s interface{}) {

	if gmp[s] {
		return
	}
	gmp[s] = true

}

func addAllInGSet(arr []interface{}) {

	for i := 0; i < len(arr); i++ {
		addInGSet(arr[i])
	}

}

func existsInGSet(s interface{}) bool {
	return gmp[s]
}

func toArrayFromGSet() []interface{} {

	keys := reflect.ValueOf(gmp).MapKeys()

	strkeys := make([]interface{}, len(keys))
	for i := 0; i < len(keys); i++ {
		strkeys[i] = keys[i].Interface()
	}

	return strkeys

}

func gsetInit() {
	gmp = make(map[interface{}]bool)
}
