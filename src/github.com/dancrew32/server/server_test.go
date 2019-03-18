package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHealthCheckHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/decode", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthCheckHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("bad status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"version":"abcdef","alive":true}`
	body := strings.TrimSpace(rr.Body.String())
	if body != expected {
		t.Errorf("body mismatch: got %v want %v", body, expected)
	}
}
