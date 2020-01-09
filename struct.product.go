package main

//PRODUCT document structure
type PRODUCT struct {
	SKU            string
	name           string
	groupID        string
	description    string
	regularPrice   float64
	promotionPrice float64
	images         []string
	searchKeywords []string
	quantity       int64
	category       []string
	color          string
	brand          string
	size           string
	active         bool
	attributes     map[string]string
	isMain         bool
	currency       string
	updated        string
}
