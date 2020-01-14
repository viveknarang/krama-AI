package main

//PRICEUPDATEREQUEST document structure
type PRICEUPDATEREQUEST struct {
	Prices map[string]PRICES `json:"Prices"`
}

//PRICES document structure
type PRICES struct {
	PromotionPrice float64 `json:"PromotionPrice"`
	RegularPrice   float64 `json:"RegularPrice"`
}
