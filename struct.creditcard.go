package main

//CREDITCARD document structure
type CREDITCARD struct {
	Name            string `json:"Name" bson:"Name" validate:"min=1,max=100"`
	CardNumber      string `json:"CardNumber" bson:"CardNumber" validate:"min=16,max=16"`
	CardExpiryMM    string `json:"CardExpiryMM" bson:"CardExpiryMM" validate:"min=1,max=2"`
	CardExpiryYY    string `json:"CardExpiryYY" bson:"CardExpiryYY" validate:"min=2,max=2"`
	SecurityCode    string `json:"SecurityCode" bson:"SecurityCode" validate:"min=3,max=3"`
	ZipCode         string `json:"ZipCode" bson:"ZipCode" validate:"min=1,max=10"`
	Default         bool   `json:"Default" bson:"Default"`
	SaveInformation bool   `json:"SaveInformation" bson:"SaveInformation"`
}
