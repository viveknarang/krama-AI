package main

//SEARCHREQUEST document structure
type SEARCHREQUEST struct {
	Query            string                                `json:"Query" bson:"Query"`
	QueryFields      []string                              `json:"QueryFields" bson:"QueryFields"`
	ResponseFields   []string                              `json:"ResponseFields" bson:"ResponseFields"`
	From             int                                   `json:"From" bson:"From"`
	To               int                                   `json:"To" bson:"To"`
	TermFacetFields  []string                              `json:"TermFacetFields" bson:"TermFacetFields"`
	RangeFacetFields []map[string][]map[string]interface{} `json:"RangeFacetFields" bson:"RangeFacetFields"`
}
