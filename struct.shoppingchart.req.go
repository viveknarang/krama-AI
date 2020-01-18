package main

//SHOPPINGCARTREQ add product request
type SHOPPINGCARTREQ struct {
	CartID     string  `json:"CartID"`
	CustomerID string  `json:"CustomerID"`
	Product    PRODUCT `json:"Product"`
	Count      int64   `json:"Count"`
	SKU        string  `json:"SKU"`
}
