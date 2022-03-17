package middleware

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
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

type spaHandler struct {
	staticPath string
	indexPath  string
}

// ServeHTTP inspects the URL path to locate a file within the static dir
// on the SPA handler. If a file is found, it will be served. If not, the
// file located at the index path on the SPA handler will be served. This
// is suitable behavior for serving an SPA (single page application).
func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// get the absolute path to prevent directory traversal
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		// if we failed to get the absolute path respond with a 400 bad request
		// and stop
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// prepend the path with the path to the static directory
	path = filepath.Join(h.staticPath, path)

	// check whether a file exists at the given path
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		// file does not exist, serve index.html
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
		// if we got an error (that wasn't that the file doesn't exist) stating the
		// file, return a 500 internal server error and stop
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// otherwise, use http.FileServer to serve the static dir
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}
