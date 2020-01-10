package main

//RESPONSE document structure
type RESPONSE struct {
	Code     int         `json:"Code"`
	Success  bool        `json:"Success"`
	Message  string      `json:"Message"`
	Time     int64       `json:"Time"`
	Response interface{} `json:"Response"`
}
