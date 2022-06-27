package middleware

import "server/models"

type goodTemperature struct {
	next itemperature
}

func (gt *goodTemperature) execute(b *models.Beer) {
	b.TemperatureStatus = allGood
	return
}

func (lt *goodTemperature) setNext(next itemperature) {
	lt.next = next
}
