package main

//ADDRESS document structure
type ADDRESS struct {
	FirstName      string `json:"FirstName"`
	LastName       string `json:"LastName"`
	AddressLineOne string `json:"AddressLineOne"`
	AddressLineTwo string `json:"AddressLineTwo"`
	City           string `json:"City"`
	State          string `json:"State"`
	Country        string `json:"Country"`
	Pincode        string `json:"Pincode"`
}
