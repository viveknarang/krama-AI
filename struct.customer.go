package main

//CUSTOMER document structure
type CUSTOMER struct {
	CustomerID string `json:"CustomerID"`
	APIKey     string `json:"APIKey"`
	Name       string `json:"Name"`
	Secret     string `json:"Secret"`
	Active     bool   `json:"Active"`
}
