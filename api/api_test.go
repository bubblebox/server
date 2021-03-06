package api_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestItemsIndex(t *testing.T) {
	request, _ := http.NewRequest("GET", "/api/v1/items", nil)
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Errorf("Expected HTTP 200 OK, but got HTTP %d instead", response.Code)
	}

	actual := response.Body.String()
	expected := `
  {
    "items": [
      {
        "code": "txt",
        "type": 1,
        "content": "Lorem ipsum",
        "created_at": "2015-11-19T12:19:33.865042825+01:00"
      },
      {
        "code": "url",
        "type": 0,
        "content": "https://ariejan.net",
        "created_at": "2015-11-19T12:19:33.865042825+01:00"
      }
    ]
  }`

	if !matchJSON(actual, expected) {
		t.Errorf("Expected %s to match %s", actual, expected)
	}
}

func TestItemShow(t *testing.T) {
	request, _ := http.NewRequest("GET", "/api/v1/items/url", nil)
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Errorf("Expected HTTP 200 OK, but got HTTP %d instead", response.Code)
	}

	actual := response.Body.String()
	expected := `
  {
    "item": {
			"code": "url",
			"type": 0,
			"content": "https://ariejan.net",
			"created_at": "2015-11-19T12:19:33.865042825+01:00"
    }
	}`

	if !matchJSON(actual, expected) {
		t.Errorf("Expected %s to match %s", actual, expected)
	}
}

func TestItemsShowNotFound(t *testing.T) {
	request, _ := http.NewRequest("GET", "/api/v1/items/unknown", nil)
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	if response.Code != http.StatusNotFound {
		t.Errorf("Expected HTTP 404 Not found, but got HTTP %d instead", response.Code)
	}
}

func TestItemCreateWithCode(t *testing.T) {
	jsonStr := []byte(`{
		"code": "teapot",
		"content": "http://google.com/teapot"}
	`)

	request, _ := http.NewRequest("POST", "/api/v1/items", bytes.NewBuffer(jsonStr))
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	if response.Code != http.StatusCreated {
		t.Errorf("Expected HTTP 201 Created, but got HTTP %d instead", response.Code)
	}

	if !database.DoesItemExist("teapot") {
		t.Error("Expeced item 'teapot' to have been persisted")
	} else {
		database.DeleteItem("teapot")
	}
}

func TestItemDestroy(t *testing.T) {
	request, _ := http.NewRequest("DELETE", "/api/v1/items/url", nil)
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Errorf("Expected HTTP 200 OK, but got HTTP %d instead", response.Code)
	}

	actual := response.Body.String()
	expected := `{}`

	if !matchJSON(actual, expected) {
		t.Errorf("Expected %s to match %s", actual, expected)
	}
}
