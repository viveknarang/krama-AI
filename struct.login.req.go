package main

//LOGIN document structure
type LOGIN struct {
	CustomerID string `json:"CustomerID", bson:"CustomerID"`
	APIKey     string `json:"APIKey", bson:"APIKey"`
}
