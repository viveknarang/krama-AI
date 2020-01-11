package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func respondWith(w http.ResponseWriter, r *http.Request, err error, message string, response interface{}, code int) {

	w.Header().Set("Content-Type", "application/json")

	var resp RESPONSE

	if err != nil {

		resp.Code = http.StatusBadRequest
		resp.Message = err.Error()
		resp.Success = false

	} else {

		resp.Code = code
		resp.Message = message
		resp.Success = true

	}

	resp.Time = time.Now().UnixNano()
	resp.Response = response

	respons, err := json.Marshal(resp)
	w.Write([]byte(respons))

}
