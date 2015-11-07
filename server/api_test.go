package main_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ariejan/firedragon/server"
)

func TestURLRedirectHandling(t *testing.T) {
	request, _ := http.NewRequest("GET", "/url", nil)
	response := httptest.NewRecorder()

	main.Router(db).ServeHTTP(response, request)

	if response.Code != 301 {
		t.Errorf("Expected HTTP 301 Redirect, but got HTTP %d instead", response.Code)
	}

	location := response.Header()["Location"][0]
	if location != "https://ariejan.net" {
		t.Errorf("Expected redirect to https://ariejan.net, got %s instead", location)
	}
}
