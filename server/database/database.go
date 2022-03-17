package database

import "server/models"

func GetAllBeers() []models.Beer {
	var beers []models.Beer

	beer1 := models.Beer{
		Id:                 "1",
		Name:               "Pilsner",
		MinimumTemperature: 4,
		MaximumTemperature: 6,
		Temperature:        0,
		TemperatureStatus:  "",
	}

	beer2 := models.Beer{
		Id:                 "2",
		Name:               "IPA",
		MinimumTemperature: 5,
		MaximumTemperature: 6,
		Temperature:        0,
		TemperatureStatus:  "",
	}

	beer3 := models.Beer{
		Id:                 "3",
		Name:               "Lager",
		MinimumTemperature: 4,
		MaximumTemperature: 7,
		Temperature:        0,
		TemperatureStatus:  "",
	}

	beer4 := models.Beer{
		Id:                 "4",
		Name:               "Stout",
		MinimumTemperature: 6,
		MaximumTemperature: 8,
		Temperature:        0,
		TemperatureStatus:  "",
	}

	beer5 := models.Beer{
		Id:                 "5",
		Name:               "Wheat beer",
		MinimumTemperature: 3,
		MaximumTemperature: 5,
		Temperature:        0,
		TemperatureStatus:  "",
	}

	beer6 := models.Beer{
		Id:                 "6",
		Name:               "Pale Ale",
		MinimumTemperature: 4,
		MaximumTemperature: 6,
		Temperature:        0,
		TemperatureStatus:  "",
	}

	beers = append(beers, beer1, beer2, beer3, beer4, beer5, beer6)

	return beers
}
