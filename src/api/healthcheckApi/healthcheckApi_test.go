package healthcheckApi

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"testing"
)

func TestGetShouldReturnExpectedResponse(t *testing.T) {
	// arrange
	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/", nil)

	dbServiceImpl := &DbService{
		checkDbConnectionStatus: func() bool {
			return true
		},
	}
	SetupDbService(dbServiceImpl)

	// act
	Get(response, request)

	// assert
	if response.Code != 200 {
		t.Fatal("OK Status Code was expected")
	}

	var dto *HealthcheckDto
	err := json.NewDecoder(response.Body).Decode(&dto)
	if err != nil || dto == nil {
		t.Fatal("Failed to Unmarshal body as HealthcheckDto")
	}
	if dto.Status != "OK" {
		t.Fatal("Expected Status to be OK")
	}
	if dto.DbStatus != "OK" {
		t.Fatal("Expected DbStatus to be OK")
	}
}

func TestGetShouldReturnExpectedResponseWhenDbStatusIsBad(t *testing.T) {
	// arrange
	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/", nil)

	dbServiceImpl := &DbService{
		checkDbConnectionStatus: func() bool {
			return false
		},
	}
	SetupDbService(dbServiceImpl)

	// act
	Get(response, request)

	// assert
	if response.Code != 200 {
		t.Fatal("OK Status Code was expected")
	}

	var dto *HealthcheckDto
	err := json.NewDecoder(response.Body).Decode(&dto)
	if err != nil || dto == nil {
		t.Fatal("Failed to Unmarshal body as HealthcheckDto")
	}
	if dto.Status != "OK" {
		t.Fatal("Expected Status to be OK")
	}
	if dto.DbStatus != "Bad" {
		t.Fatal("Expected DbStatus to be Bad")
	}
}
