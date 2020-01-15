package main

//PLATFORMCUSTOMER document structure
type PLATFORMCUSTOMER struct {
	CustomerID string `json:"CustomerID"`
	APIKey     string `json:"APIKey"`
	Name       string `json:"Name"`
	Secret     string `json:"Secret"`
	Active     bool   `json:"Active"`
}
