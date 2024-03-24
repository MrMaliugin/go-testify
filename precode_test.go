package main

import (
	"github.com/stretchr/testify/require"
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

	require.Equal(t, http.StatusOK, responseRecorder.Code, "expected status code")
	require.NotEmpty(t, responseRecorder.Body.String(), "response body is empty")
}

func TestMainHandlerWhenMissingCityMsc(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=10&city=UnExistsCity", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusBadRequest, responseRecorder.Code, "expected status code")

	expected := `wrong city value`

	require.Equal(t, expected, responseRecorder.Body.String(), "the answer is correct")

}
func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code, "expected status code")

	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")

	require.Len(t, list, totalCount, "input error, fewer cafes specified")
}
