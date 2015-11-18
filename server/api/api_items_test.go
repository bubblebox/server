package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestItemView(t *testing.T) {
	request, _ := http.NewRequest("GET", "/api/v1/items/url", nil)
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != 200 {
		t.Errorf("Expected HTTP 200 OK, but got HTTP %d instead", response.Code)
	}

	actual := response.Body.String()
	expected := `
  {
    "data": {
      "attributes": {
        "content": "https://ariejan.net",
        "type": 0
      },
      "id": "url",
      "type": "items"
    },
    "meta": {
      "generator": "Fire Dragon"
    }
  }`

	if !matchJSON(actual, expected) {
		t.Errorf("Expected %s to match %s", actual, expected)
	}
}

func TestItemList(t *testing.T) {
	request, _ := http.NewRequest("GET", "/api/v1/items", nil)
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != 200 {
		t.Errorf("Expected HTTP 200 OK, but got HTTP %d instead", response.Code)
	}

	actual := response.Body.String()
	expected := `
  {
    "data": [
			{
				"attributes": {
					"content": "Lorem ipsum",
					"type": 1
				},
				"id": "txt",
				"type": "items"
			},
			{
				"attributes": {
					"content": "https://ariejan.net",
					"type":0
				},
				"id": "url",
				"type": "items"
			}
		],
    "meta": {
      "generator": "Fire Dragon"
    }
  }`

	if !matchJSON(actual, expected) {
		t.Errorf("Expected %s to match %s", actual, expected)
	}
}

func TestItemViewNotFound(t *testing.T) {
	request, _ := http.NewRequest("GET", "/api/v1/items/notfound", nil)
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != 404 {
		t.Errorf("Expected HTTP 404 Not Found, but got HTTP %d instead", response.Code)
	}

	actual := response.Body.String()
	expected := `
  {
    "errors": [
      {
        "status": "404",
        "title": "Item with ID 'notfound' not found"
      }
    ]
  }`

	if !matchJSON(actual, expected) {
		t.Errorf("Expected %s to match %s", actual, expected)
	}
}
