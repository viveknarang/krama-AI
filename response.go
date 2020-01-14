package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func respondWith(w http.ResponseWriter, r *http.Request, err error, message string, response interface{}, code int, success bool) {

	w.Header().Set("Content-Type", "application/json")

	var resp RESPONSE

	if err != nil {

		resp.Message = err.Error()

	} else {

		resp.Message = message

	}

	resp.Code = code
	resp.Success = success
	resp.Response = response
	resp.Time = time.Now().UnixNano()

	respons, err := json.Marshal(resp)
	w.Write([]byte(respons))

}
