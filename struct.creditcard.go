package main

//CREDITCARD document structure
type CREDITCARD struct {
	Name         string `json:"Name"`
	CardNumber   string `json:"CardNumber"`
	CardExpiryMM string `json:"CardExpiryMM"`
	CardExpiryYY string `json:"CardExpiryYY"`
	SecurityCode string `json:"SecurityCode"`
	ZipCode      string `json:"ZipCode"`
}
