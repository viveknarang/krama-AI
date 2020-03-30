package main

import "testing"

func TestURLValidMethod(t *testing.T) {

	t.Log("Testing the isValidURL function ...")

	validURL := isValidURL("https://homepages.cae.wisc.edu/~ece533/images/boat.png")

	if !validURL {
		t.Errorf("isValidURL function not working properly. A valid URL is returning false instead of true")
	}

	invalidURL := isValidURL(" DDSSD ..... https://homepages.cae.wisc.edu/~ece533/images/boat.png .....")

	if invalidURL {
		t.Errorf("isValidURL function not working properly. An invalid URL is returning true instead of false")
	}

}
