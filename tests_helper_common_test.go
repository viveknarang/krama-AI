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

func TestContainsInArray(t *testing.T) {

	t.Log("Testing the containsInArray function ...")

	var a []string
	a = append(a, "A")
	a = append(a, "B")

	if !containsInArray(a, "A") {
		t.Errorf("containsInArray:: Looking for an existing value in the array is returning false")
	}

	if containsInArray(a, "C") {
		t.Errorf("containsInArray:: Looking for a non-existant value in an array is returning true")
	}

}

func TestRemoveElementsFromArray(t *testing.T) {

	t.Log("Testing the removeElementsFromArray function ...")

	var a []string
	a = append(a, "A")
	a = append(a, "B")
	var b []string
	b = append(b, "B")

	var d []string
	d = append(d, "X")

	c := removeElementsFromArray(a, b)

	if len(c) == 2 || (len(c) == 2 && c[1] == "B") {
		t.Errorf("removeElementsFromArray:: Trying to remove an existing element. The element still exists.")
	}

	e := removeElementsFromArray(a, d)

	if len(e) != 2 {
		t.Errorf("removeElementsFromArray:: Trying to remove a non-existant value in an array is potentially removing some other existing value")
	}

}

func TestCleanCategoryPath(t *testing.T) {

	t.Log("Testing the cleanCategoryPath function ...")

	if cleanCategoryPath("> A >> B") != "A>B" {
		t.Errorf("cleanCategoryPath:: Testing `> A >> B` is not returning `A>B`")
	}

	if cleanCategoryPath("A >> B") != "A>B" {
		t.Errorf("cleanCategoryPath:: Testing `A >> B` is not returning `A>B`")
	}

	if cleanCategoryPath("> A >> B >>") != "A>B" {
		t.Errorf("cleanCategoryPath:: Testing `> A >> B >>` is not returning `A>B`")
	}

	if cleanCategoryPath("> A >>") != "A" {
		t.Errorf("cleanCategoryPath:: Testing `> A >>` is not returning `A`")
	}

	if cleanCategoryPath("A    >      B") != "A>B" {
		t.Errorf("cleanCategoryPath:: Testing `A    >      B` is not returning `A>B`")
	}

	if cleanCategoryPath("         A       >      B        ") != "A>B" {
		t.Errorf("cleanCategoryPath:: Testing `         A       >      B        ` is not returning `A>B`")
	}

}
