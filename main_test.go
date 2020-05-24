package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"fmt"
	"strings"
)

func TestHealthCheck(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)

	if err != nil {
		t.Errorf("An error occured. %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(greeting)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code. Got %v, want %v", status, http.StatusOK)
	}

	dt := time.Now()
	expected := fmt.Sprintf("%s", dt.Format("02-01-2006 15:04:05 Monday"))
	if ! strings.Contains(rr.Body.String(), expected) {
		t.Errorf("Handler did not contain correct time. Got %v, want %v", rr.Body.String(), expected)
	}
}
