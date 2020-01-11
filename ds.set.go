package main

import "reflect"

var mp = make(map[string]bool)

func addInSet(s string) {

	if mp[s] {
		return
	}

	mp[s] = true

}

func addAllInSet(arr []string) {

	for i := 0; i < len(arr); i++ {
		addInSet(arr[i])
	}

}

func existsInSet(s string) bool {

	return mp[s]

}

func toArrayFromSet() []string {

	keys := reflect.ValueOf(mp).MapKeys()

	strkeys := make([]string, len(keys))
	for i := 0; i < len(keys); i++ {
		strkeys[i] = keys[i].String()
	}

	return strkeys

}

func setInit() {

	mp = make(map[string]bool)

}
