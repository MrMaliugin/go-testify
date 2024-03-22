package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMainHandlerWhenOk(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=4&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("expected status code: %d, got %d", http.StatusOK, status)
	}
}

func TestMainHandlerWhenMissingCityMsc(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=10&city=novgorod", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusBadRequest {
		t.Errorf("expected status code: %d, got %d", http.StatusBadRequest, status)
	}

	expected := `wrong city value`
	if responseRecorder.Body.String() != expected {
		t.Errorf("expected body: %s, got %s", expected, responseRecorder.Body.String())
	}
}
func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusOK {
		t.Fatalf("expected status code: %d, got %d", http.StatusOK, status)
	}

	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")

	if len(list) != totalCount {
		t.Errorf("expected cafe count: %d, got %d", totalCount, len(list))
	}
}
