package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/romana/rlog"
)

func respondWith(w http.ResponseWriter, r *http.Request, err error, message string, response interface{}, code int, success bool) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	var resp RESPONSE

	if err != nil {

		rlog.Debug("respondWith() with code " + strconv.Itoa(code) + " for: " + r.Method + " : " + r.URL.Path + " has Error: " + err.Error())
		resp.Message = err.Error()

	} else {

		rlog.Debug("respondWith() with code " + strconv.Itoa(code) + " for: " + r.Method + " : " + r.URL.Path)
		resp.Message = message

	}

	resp.Code = code
	resp.Success = success
	resp.Response = response
	resp.Time = time.Now().UnixNano()

	respons, err := json.Marshal(resp)
	w.Write([]byte(respons))

}
