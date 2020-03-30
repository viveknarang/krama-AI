package main

//RESPONSE document structure
type RESPONSE struct {
	Code     int         `json:"Code" bson:"Code"`
	Success  bool        `json:"Success" bson:"Success"`
	Message  string      `json:"Message" bson:"Message"`
	Time     int64       `json:"Time" bson:"Time"`
	Response interface{} `json:"Response" bson:"Response"`
}
