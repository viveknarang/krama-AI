package main

//CREDITCARD document structure
type CREDITCARD struct {
	Name            string `json:"Name" bson:"Name"`
	CardNumber      string `json:"CardNumber" bson:"CardNumber"`
	CardExpiryMM    string `json:"CardExpiryMM" bson:"CardExpiryMM"`
	CardExpiryYY    string `json:"CardExpiryYY" bson:"CardExpiryYY"`
	SecurityCode    string `json:"SecurityCode" bson:"SecurityCode"`
	ZipCode         string `json:"ZipCode" bson:"ZipCode"`
	Default         bool   `json:"Default" bson:"Default"`
	SaveInformation bool   `json:"SaveInformation" bson:"SaveInformation"`
}
