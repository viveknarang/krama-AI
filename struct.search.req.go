package main

//SEARCHREQUEST document structure
type SEARCHREQUEST struct {
	Q      string   `json:"Q"`
	Fields []string `json:"Fields"`
}
