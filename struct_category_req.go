package main

//CATEGORYREQUEST document structure
type CATEGORYREQUEST struct {
	Path     string `json:"Path" bson:"Path" validate:"min=1,max=1024"`
	Category string `json:"Category" bson:"Category" validate:"min=1,max=100"`
}
