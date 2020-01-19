package main

//PRODREVIEWREQ document structure
type PRODREVIEWREQ struct {
	From      int64  `json:"From" bson:"From"`
	To        int64  `json:"To" bson:"To"`
	Order     int    `json:"Order" bson:"Order"`
	SortField string `json:"SortField" bson:"SortField"`
}
