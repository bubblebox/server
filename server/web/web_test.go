package web_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestItemNotFound(t *testing.T) {
	request, _ := http.NewRequest("GET", "/nope", nil)
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	if response.Code != http.StatusNotFound {
		t.Errorf("Expected HTTP 404 Not Found, but got HTTP %d instead", response.Code)
	}
}

func TestURLRedirect(t *testing.T) {
	request, _ := http.NewRequest("GET", "/url", nil)
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	if response.Code != http.StatusMovedPermanently {
		t.Errorf("Expected HTTP 301 Moved Permanently, but got HTTP %d instead", response.Code)
	}

	header := response.Header().Get("Location")
	expectedURL := "https://ariejan.net"
	if header != expectedURL {
		t.Errorf("Expected redirect to 'https://ariejan.net', got '%s' instead.", header)
	}

	body := response.Body.String()
	expectedTag := "<a href=\"https://ariejan.net\">Moved Permanently</a>"
	if !strings.Contains(body, expectedTag) {
		t.Errorf("Expected redirect body to contain '%s', but it did not.", expectedTag)
	}
}
