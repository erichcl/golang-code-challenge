package middleware

import "server/models"

type highTemperature struct {
	next itemperature
}

func (ht *highTemperature) execute(b *models.Beer) {
	if b.Temperature > b.MaximumTemperature {
		b.TemperatureStatus = tooHigh
		return
	} else {
		ht.next.execute(b)
	}
}

func (lt *highTemperature) setNext(next itemperature) {
	lt.next = next
}
