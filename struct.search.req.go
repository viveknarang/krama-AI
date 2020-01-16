package main

//SEARCHREQUEST document structure
type SEARCHREQUEST struct {
	Query          string   `json:"Query"`
	QueryFields    []string `json:"QueryFields"`
	ResponseFields []string `json:"ResponseFields"`
	From           int      `json:"From"`
	To             int      `json:"To"`
	FacetFields    []string `json:"FacetFields"`
}
