package main

//CATEGORYTREENODE document structure
type CATEGORYTREENODE struct {
	CategoryID string   `json:"CategoryID" bson:"CategoryID" validate:"min=1,max=100,hasNoSpaces"`
	Name       string   `json:"Name" bson:"Name" validate:"min=1,max=100"`
	Parent     string   `json:"Parent" bson:"Parent" validate:"min=1,max=100"`
	Children   []string `json:"Children" bson:"Children" validate:"min=1,max=100"`
	SKUs       []string `json:"SKUs" bson:"SKUs" validate:"min=1,max=100,hasNoSpaces"`
}
