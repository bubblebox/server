package main_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ariejan/firedragon/server"
	"github.com/shwoodard/jsonapi"
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

func TestItemViewHandling(t *testing.T) {
	request, _ := http.NewRequest("GET", "/api/v1/items/url", nil)
	response := httptest.NewRecorder()

	main.Router(db).ServeHTTP(response, request)

	if response.Code != 200 {
		t.Errorf("Expected HTTP 200 OK, but got HTTP %d instead", response.Code)
	}

	item := &main.Item{}
	if err := jsonapi.UnmarshalPayload(response.Body, item); err != nil {
		t.Error("Could not unmarshal returned JSON")
	}

	if item.Code != "url" {
		t.Error("Code does not match requested item code.")
	}
	if item.Content != "https://ariejan.net" {
		t.Error("Content does not match expected item content.")
	}
}

func TestItemNotFoundViewHandling(t *testing.T) {
	request, _ := http.NewRequest("GET", "/api/v1/items/nope", nil)
	response := httptest.NewRecorder()

	main.Router(db).ServeHTTP(response, request)

	if response.Code != 404 {
		t.Errorf("Expected HTTP 404 Not Found, but got HTTP %d instead", response.Code)
	}
}
