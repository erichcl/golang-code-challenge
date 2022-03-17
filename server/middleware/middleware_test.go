package middleware

import (
	"net/http"
	"net/http/httptest"
	"server/models"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSetBeerTempStatus(t *testing.T) {

	mockBeer := models.Beer{
		Id:                 "1",
		Name:               "Pilsner",
		MinimumTemperature: 4,
		MaximumTemperature: 6,
		Temperature:        0,
		TemperatureStatus:  "",
	}

	SetBeerTempStatus(&mockBeer)
	assert.Equal(t, mockBeer.TemperatureStatus, tooLow)

	mockBeer.Temperature = 5
	SetBeerTempStatus(&mockBeer)
	assert.Equal(t, mockBeer.TemperatureStatus, allGood)

	mockBeer.Temperature = 10
	SetBeerTempStatus(&mockBeer)
	assert.Equal(t, mockBeer.TemperatureStatus, tooHigh)
}

func TestGetAllProducts(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/products", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetAllProducts)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusOK)
}
