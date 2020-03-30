package main

//SHOPPINGCART document structure
type SHOPPINGCART struct {
	CartID        string             `json:"CartID" bson:"CartID"`
	CustomerID    string             `json:"CustomerID" bson:"CustomerID"`
	ProductsCount map[string]int64   `json:"ProductsCount" bson:"ProductsCount"`
	Products      map[string]PRODUCT `json:"Products" bson:"Products"`
	Total         float64            `json:"Total" bson:"Total"`
	Currency      string             `json:"Currency" bson:"Currency"`
	Updated       int64              `json:"Updated" bson:"Updated"`
}
