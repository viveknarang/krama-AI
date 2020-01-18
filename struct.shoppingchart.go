package main

//SHOPPINGCART document structure
type SHOPPINGCART struct {
	CartID        string             `json:"CartID"`
	CustomerID    string             `json:"CustomerID"`
	ProductsCount map[string]int64   `json:"ProductsCount"`
	Products      map[string]PRODUCT `json:"Products"`
	Total         float64            `json:"Total"`
	Currency      string             `json:"Currency"`
	Updated       int64              `json:"Updated"`
}
