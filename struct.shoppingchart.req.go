package main

//SHOPPINGCARTREQ add product request
type SHOPPINGCARTREQ struct {
	CartID     string  `json:"CartID" bson:"CartID"`
	CustomerID string  `json:"CustomerID" bson:"CustomerID"`
	Product    PRODUCT `json:"Product" bson:"Product"`
	Count      int64   `json:"Count" bson:"Count"`
	SKU        string  `json:"SKU" bson:"SKU"`
}
