package middleware

import (
	"encoding/json"
	"net/http"
	"server/database"
	"server/models"
)

const tooLow string = "too low"
const tooHigh string = "too high"
const allGood string = "all good"

func SetBeerTempStatus(b *models.Beer) {
	if b.Temperature < b.MinimumTemperature {
		b.TemperatureStatus = tooLow
	}

	if b.Temperature > b.MaximumTemperature {
		b.TemperatureStatus = tooHigh
	}

	if b.Temperature >= b.MinimumTemperature && b.Temperature <= b.MaximumTemperature {
		b.TemperatureStatus = allGood
	}
}

func GetAllProducts(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	beers := database.GetAllBeers()

	resultChannel := make(chan models.ResultsItem)

	for _, beer := range beers {
		go SendGetSensorAsync(beer.Id, resultChannel)
	}

	result := make(map[string]int8)

	for range beers {
		item := <-resultChannel
		result[item.Id] = item.Res.Temperature
	}

	for j, beer := range beers {
		beers[j].Temperature = result[beer.Id]
		SetBeerTempStatus(&beers[j])
	}

	json.NewEncoder(w).Encode(beers)
}
