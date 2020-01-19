package main

//PRODUCTGROUP document structure
type PRODUCTGROUP struct {
	GroupID               string                   `json:"GroupID" bson:"GroupID"`
	Name                  string                   `json:"Name" bson:"Name"`
	Description           string                   `json:"Description" bson:"Description"`
	RegularPriceMin       float64                  `json:"RegularPriceMin" bson:"RegularPriceMin"`
	RegularPriceMax       float64                  `json:"RegularPriceMax" bson:"RegularPriceMax"`
	PromotionPriceMin     float64                  `json:"PromotionPriceMin" bson:"PromotionPriceMin"`
	PromotionPriceMax     float64                  `json:"PromotionPriceMax" bson:"PromotionPriceMax"`
	Skus                  []string                 `json:"Skus" bson:"Skus"`
	Images                []string                 `json:"Images" bson:"Images"`
	SearchKeywords        []string                 `json:"SearchKeywords" bson:"SearchKeywords"`
	Category              []string                 `json:"Category" bson:"Category"`
	Colors                []string                 `json:"Colors" bson:"Colors"`
	Brands                []string                 `json:"Brands" bson:"Brands"`
	Sizes                 []string                 `json:"Sizes" bson:"Sizes"`
	Active                bool                     `json:"Active" bson:"Active"`
	Currency              string                   `json:"Currency" bson:"Currency"`
	Updated               int64                    `json:"Updated" bson:"Updated"`
	Products              map[string]PRODUCT       `json:"Products" bson:"Products"`
	Attributes            map[string][]interface{} `json:"Attributes" bson:"Attributes"`
	CumulativeReviewStars int                      `json:"CumulativeReviewStars" bson:"CumulativeReviewStars"`
	CumulativeReviewCount int64                    `json:"CumulativeReviewCount" bson:"CumulativeReviewCount"`
}
