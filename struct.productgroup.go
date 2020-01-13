package main

//PRODUCTGROUP document structure
type PRODUCTGROUP struct {
	GroupID           string                 `json:"GroupID"`
	Name              string                 `json:"Name"`
	Description       string                 `json:"Description"`
	RegularPriceMin   float64                `json:"RegularPriceMin"`
	RegularPriceMax   float64                `json:"RegularPriceMax"`
	PromotionPriceMin float64                `json:"PromotionPriceMin"`
	PromotionPriceMax float64                `json:"PromotionPriceMax"`
	Skus              []string               `json:"Skus"`
	Images            []string               `json:"Images"`
	SearchKeywords    []string               `json:"SearchKeywords"`
	Category          []string               `json:"Category"`
	Colors            []string               `json:"Colors"`
	Brands            []string               `json:"Brands"`
	Sizes             []string               `json:"Sizes"`
	Active            bool                   `json:"Active"`
	Currency          string                 `json:"Currency"`
	Updated           int64                  `json:"Updated"`
	Products          map[string]PRODUCT     `json:"Products"`
	Attributes        map[string]interface{} `json:"Attributes"`
}
