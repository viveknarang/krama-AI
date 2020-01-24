package main

import (
	"reflect"
)

func addInSet(s string, mp map[string]bool) {

	if mp[s] {
		return
	}
	mp[s] = true

}

func addAllInSet(arr []string, mp map[string]bool) {

	for i := 0; i < len(arr); i++ {
		addInSet(arr[i], mp)
	}

}

func existsInSet(s string, mp map[string]bool) bool {
	return mp[s]
}

func toArrayFromSet(mp map[string]bool) []string {

	keys := reflect.ValueOf(mp).MapKeys()

	strkeys := make([]string, len(keys))
	for i := 0; i < len(keys); i++ {
		strkeys[i] = keys[i].String()
	}

	return strkeys

}
