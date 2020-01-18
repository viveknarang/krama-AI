package main

//ADDRESS document structure
type ADDRESS struct {
	FirstName      string `json:"FirstName" bson:"FirstName"`
	LastName       string `json:"LastName" bson:"LastName"`
	AddressLineOne string `json:"AddressLineOne" bson:"AddressLineOne"`
	AddressLineTwo string `json:"AddressLineTwo" bson:"AddressLineTwo"`
	City           string `json:"City" bson:"City"`
	State          string `json:"State" bson:"State"`
	Country        string `json:"Country" bson:"Country"`
	Pincode        string `json:"Pincode" bson:"Pincode"`
	Default        bool   `json:"Default" bson:"Default"`
}
