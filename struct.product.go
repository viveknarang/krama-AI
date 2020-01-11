package main

//PRODUCT document structure
type PRODUCT struct {
	Sku            string      `json:"Sku"`
	Name           string      `json:"Name"`
	GroupID        string      `json:"GroupID"`
	Description    string      `json:"Description"`
	RegularPrice   float64     `json:"RegularPrice"`
	PromotionPrice float64     `json:"PromotionPrice"`
	Images         []string    `json:"Images"`
	SearchKeywords []string    `json:"SearchKeywords"`
	Quantity       int64       `json:"Quantity"`
	Category       []string    `json:"Category"`
	Color          string      `json:"Color"`
	Brand          string      `json:"Brand"`
	Size           string      `json:"Size"`
	Active         bool        `json:"Active"`
	Attributes     interface{} `json:"Attributes"`
	IsMain         bool        `json:"IsMain"`
	Currency       string      `json:"Currency"`
	Updated        string      `json:"Updated"`
}
