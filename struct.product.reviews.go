package main

//PRODUCTREVIEW document structure
type PRODUCTREVIEW struct {
	ReviewID    string  `json:"ReviewID" bson:"ReviewID" validate:"min=1,max=100"`
	Time        int64   `json:"Time" bson:"Time"`
	GroupID     string  `json:"GroupID" bson:"GroupID" validate:"min=1,max=100"`
	CustomerID  string  `json:"CustomerID" bson:"CustomerID" validate:"min=1,max=100"`
	Stars       float64 `json:"Stars" bson:"Stars" validate:"validStarRating"`
	Description string  `json:"Description" bson:"Description" validate:"min=1,max=10240"`
}
