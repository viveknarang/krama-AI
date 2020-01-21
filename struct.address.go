package main

//ADDRESS document structure
type ADDRESS struct {
	FirstName      string `json:"FirstName" bson:"FirstName" validate:"min=1,max=100"`
	LastName       string `json:"LastName" bson:"LastName" validate:"min=1,max=100"`
	AddressLineOne string `json:"AddressLineOne" bson:"AddressLineOne" validate:"min=1,max=200"`
	AddressLineTwo string `json:"AddressLineTwo" bson:"AddressLineTwo" validate:"min=1,max=200"`
	City           string `json:"City" bson:"City" validate:"min=1,max=100"`
	State          string `json:"State" bson:"State" validate:"min=1,max=100"`
	Country        string `json:"Country" bson:"Country" validate:"min=1,max=100"`
	Pincode        string `json:"Pincode" bson:"Pincode" validate:"min=1,max=10"`
	Default        bool   `json:"Default" bson:"Default"`
}
