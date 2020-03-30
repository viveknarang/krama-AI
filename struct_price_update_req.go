package main

//PRICEUPDATEREQUEST document structure
type PRICEUPDATEREQUEST struct {
	Prices map[string]PRICES `json:"Prices" bson:"Prices"`
}

//PRICES document structure
type PRICES struct {
	PromotionPrice float64 `json:"PromotionPrice" bson:"PromotionPrice"`
	RegularPrice   float64 `json:"RegularPrice" bson:"RegularPrice"`
}
