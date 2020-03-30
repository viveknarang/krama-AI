package main

//INVENTORY document structure
type INVENTORY struct {
	Sku      string `json:"Sku" bson:"Sku"`
	Quantity int64  `json:"Quantity" bson:"Quantity"`
	Updated  int64  `json:"Updated" bson:"Updated"`
}
