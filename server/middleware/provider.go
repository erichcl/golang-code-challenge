package middleware

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"server/models"
)

const apiUrlRequest string = "https://temperature-sensor-service.herokuapp.com/sensor/"

func SendGetSensorAsync(id string, rc chan models.ResultsItem) {

	response, err := http.Get(apiUrlRequest + id)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var sensor models.Sensor
	if err := json.Unmarshal(body, &sensor); err != nil {
		panic(err)
	}
	rc <- models.ResultsItem{Id: id, Res: sensor}
}
