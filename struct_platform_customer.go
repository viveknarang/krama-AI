package main

//PLATFORMCUSTOMER document structure
type PLATFORMCUSTOMER struct {
	CustomerID string `json:"CustomerID" bson:"CustomerID"`
	APIKey     string `json:"APIKey" bson:"APIKey"`
	Name       string `json:"Name" bson:"Name"`
	Secret     string `json:"Secret" bson:"Secret"`
	Active     bool   `json:"Active" bson:"Active"`
}
