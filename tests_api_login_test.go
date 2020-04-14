package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLogin(t *testing.T) {

	t.Log("Testing the Login Endpoint ...")

	var jsonStr = []byte(`{"CustomerID":"6476154099","APIKey":"zaCELgL.0imfnc8mVLWwsAawjYr4Rx-Af50DDqtlx"}`)

	req, err := http.NewRequest("POST", "localhost:9005/customers/v1/login", bytes.NewBuffer(jsonStr))

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(login)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := ``
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

}
