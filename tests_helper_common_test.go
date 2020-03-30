package main

import "testing"

func TestIsValidURL(t *testing.T) {

	t.Log("Testing the isValidURL function ...")

	validURL := isValidURL("https://homepages.cae.wisc.edu/~ece533/images/boat.png")

	if !validURL {
		t.Errorf("isValidURL:: A valid URL is returning false instead of true")
	}

	invalidURL := isValidURL(" DDSSD ..... https://homepages.cae.wisc.edu/~ece533/images/boat.png .....")

	if invalidURL {
		t.Errorf("isValidURL:: An invalid URL is returning true instead of false")
	}

}

func TestIsValidEmail(t *testing.T) {

	t.Log("Testing the isValidEmail function ...")

	validEmail := isValidEmail("donald@mentallyill.com")

	if !validEmail {
		t.Errorf("isValidEmail:: A valid email address is retuning false")
	}

	invalidEmail := isValidEmail("donald#mentallyill.com")

	if invalidEmail {
		t.Errorf("isValidEmail:: An invalid email address is returning true")
	}

}

func TestIsValidJSON(t *testing.T) {

	t.Log("Testing the isValidJSON function ...")

	validJSON := isValidJSON("{ \"a\":\"b\" }")

	if !validJSON {
		t.Errorf("isValidJSON:: A valid JSON is returning false")
	}

	invalidJSON := isValidJSON("{ kjkldjfkldjdkljd $#VFGGGF }")

	if invalidJSON {
		t.Errorf("isValidJSON:: An invalid JSON is returning true")
	}

}
