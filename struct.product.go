package main

//PRODUCT document structure
type PRODUCT struct {
	Sku            string                 `json:"Sku" bson:"Sku"`
	Name           string                 `json:"Name" bson:"Name"`
	GroupID        string                 `json:"GroupID" bson:"GroupID"`
	Description    string                 `json:"Description" bson:"Description"`
	RegularPrice   float64                `json:"RegularPrice" bson:"RegularPrice"`
	PromotionPrice float64                `json:"PromotionPrice" bson:"PromotionPrice"`
	Images         []string               `json:"Images" bson:"Images"`
	SearchKeywords []string               `json:"SearchKeywords" bson:"SearchKeywords"`
	Quantity       int64                  `json:"Quantity" bson:"-"`
	Category       []string               `json:"Category" bson:"Category"`
	Color          string                 `json:"Color" bson:"Color"`
	Brand          string                 `json:"Brand" bson:"Brand"`
	Size           string                 `json:"Size" bson:"Size"`
	Active         bool                   `json:"Active" bson:"Active"`
	Attributes     map[string]interface{} `json:"Attributes" bson:"Attributes"`
	IsMain         bool                   `json:"IsMain" bson:"IsMain"`
	Currency       string                 `json:"Currency" bson:"Currency"`
	Updated        int64                  `json:"Updated" bson:"Updated"`
}
