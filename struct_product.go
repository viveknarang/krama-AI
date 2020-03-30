package main

//PRODUCT document structure
type PRODUCT struct {
	Sku                   string                 `json:"Sku" bson:"Sku" validate:"min=1,max=50,hasNoSpaces"`
	Name                  string                 `json:"Name" bson:"Name" validate:"min=1,max=100"`
	GroupID               string                 `json:"GroupID" bson:"GroupID" validate:"min=1,max=100,hasNoSpaces"`
	Description           string                 `json:"Description" bson:"Description" validate:"max=1024"`
	RegularPrice          float64                `json:"RegularPrice" bson:"RegularPrice" validate:"min=0.0, checkMaxFloat"`
	PromotionPrice        float64                `json:"PromotionPrice" bson:"PromotionPrice" validate:"min=0.0, checkMaxFloat"`
	Images                []string               `json:"Images" bson:"Images" validate:"size=100"`
	SearchKeywords        []string               `json:"SearchKeywords" bson:"SearchKeywords" validate:"size=100"`
	Quantity              int64                  `json:"Quantity" bson:"-" validate:"min=0"`
	Category              []string               `json:"Category" bson:"Category"`
	Color                 string                 `json:"Color" bson:"Color" validate:"size=100"`
	Brand                 string                 `json:"Brand" bson:"Brand" validate:"size=100"`
	Size                  string                 `json:"Size" bson:"Size" validate:"size=100"`
	Active                bool                   `json:"Active" bson:"Active"`
	Selectors             map[string]interface{} `json:"Selectors" bson:"Selectors" validate:"size=500"`
	Attributes            map[string]interface{} `json:"Attributes" bson:"Attributes" validate:"size=500"`
	IsMain                bool                   `json:"IsMain" bson:"IsMain"`
	Currency              string                 `json:"Currency" bson:"Currency" validate:"isValidCurrency"`
	ComplementaryProducts []string               `json:"ComplementaryProducts" bson:"ComplementaryProducts" validate:"size=100"`
	Updated               int64                  `json:"Updated" bson:"Updated"`
}
