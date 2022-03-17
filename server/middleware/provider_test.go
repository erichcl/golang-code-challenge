package middleware

import (
	"fmt"
	"server/models"
	"testing"

	"github.com/go-playground/assert/v2"
	gock "gopkg.in/h2non/gock.v1"
)

type mockInterface struct {
	mainChannel chan models.ResultsItem
}

func TestSendGetSensorAsync(t *testing.T) {
	defer gock.Off()

	mock := mockInterface{
		mainChannel: make(chan models.ResultsItem),
	}
	defer close(mock.mainChannel)

	gock.New("https://temperature-sensor-service.herokuapp.com/sensor/").
		Get("1").
		Reply(200).
		JSON(models.Sensor{Id: "1", Temperature: 8})

	go SendGetSensorAsync("1", mock.mainChannel)
	result := <-mock.mainChannel
	expectedRes := models.ResultsItem{Id: "1", Res: models.Sensor{Id: "1", Temperature: 8}}

	assert.NotEqual(t, nil, result)
	assert.Equal(t, expectedRes, result)
}

func TestSendGetSensorAsyncFail(t *testing.T) {
	defer gock.Off()
	mock := mockInterface{
		mainChannel: make(chan models.ResultsItem),
	}

	gock.New("https://temperature-sensor-service.herokuapp.com/sensor/").
		Get("1").
		ReplyError(fmt.Errorf("error"))

	go func() {
		defer close(mock.mainChannel)
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected panic")
			}
		}()
		SendGetSensorAsync("1", mock.mainChannel)
	}()

	result := <-mock.mainChannel
	assert.NotEqual(t, nil, result)
}
