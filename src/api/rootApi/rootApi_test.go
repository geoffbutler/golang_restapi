package rootApi

import (
	"net/http"
	"net/http/httptest"

	"testing"
)

func TestGetShouldReturnOkStatusCode(t *testing.T) {
	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/", nil)

	Get(response, request)

	if response.Code != 200 {
		t.Fatal("OK Status Code was expected")
	}
}
