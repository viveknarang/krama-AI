package main

//PRODUCTREVIEW document structure
type PRODUCTREVIEW struct {
	ReviewID    string  `json:"ReviewID" bson:"ReviewID"`
	Time        int64   `json:"Time" bson:"Time"`
	GroupID     string  `json:"GroupID" bson:"GroupID"`
	CustomerID  string  `json:"CustomerID" bson:"CustomerID"`
	Stars       float64 `json:"Stars" bson:"Stars"`
	Description string  `json:"Description" bson:"Description"`
}
