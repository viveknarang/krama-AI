package main

//INVENTORYUPDATEREQUEST document structure
type INVENTORYUPDATEREQUEST struct {
	Quantity map[string]int64 `json:"Quantity" bson:"Quantity"`
}
