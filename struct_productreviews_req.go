package main

//PRODREVIEWREQ document structure
type PRODREVIEWREQ struct {
	From      int64  `json:"From" bson:"From"`
	To        int64  `json:"To" bson:"To"`
	Order     int    `json:"Order" bson:"Order" validate:"isValidSortOrder"`
	SortField string `json:"SortField" bson:"SortField" validate:"min=1,max=100"`
}
