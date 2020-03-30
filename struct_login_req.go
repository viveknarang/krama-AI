package main

//LOGIN document structure
type LOGIN struct {
	CustomerID string `json:"CustomerID" bson:"CustomerID" validate:"min=1,max=100,hasNoSpaces"`
	APIKey     string `json:"APIKey" bson:"APIKey" validate:"min=1,max=10240,hasNoSpaces"`
}
