package main

//SIRQ document structure
type SIRQ struct {
	CategoryPath string `json:"CategoryPath" bson:"CategoryPath" validate:"min=1,max=1024"`
}
