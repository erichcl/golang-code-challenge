package middleware

import "server/models"

type lowTemperature struct {
	next itemperature
}

func (lt *lowTemperature) execute(b *models.Beer) {
	if b.Temperature < b.MinimumTemperature {
		b.TemperatureStatus = tooLow
		return
	} else {
		lt.next.execute(b)
	}
}

func (lt *lowTemperature) setNext(next itemperature) {
	lt.next = next
}
