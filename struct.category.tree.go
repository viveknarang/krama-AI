package main

//CATEGORYTREENODE document structure
type CATEGORYTREENODE struct {
	CategoryID     string   `json:"CategoryID" bson:"CategoryID" validate:"min=1,max=100,hasNoSpaces"`
	CategoryName   string   `json:"CategoryName" bson:"CategoryName" validate:"min=1,max=100"`
	ParentCategory string   `json:"Parent" bson:"Parent" validate:"min=1,max=100"`
	ChildCategory  []string `json:"ChildCategory" bson:"ChildCategory" validate:"min=1,max=100"`
	SKUs           []string `json:"SKUs" bson:"SKUs" validate:"min=1,max=100,hasNoSpaces"`
}
