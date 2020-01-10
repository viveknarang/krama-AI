package main

//RESPONSE document structure
type RESPONSE struct {
	Code     string      `json:"Code"`
	Success  bool        `json:"Success"`
	Message  string      `json:"Message"`
	Response interface{} `json:"Response"`
	Time     int64       `json:"Time"`
}
