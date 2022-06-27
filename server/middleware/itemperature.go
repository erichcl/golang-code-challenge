package middleware

import "server/models"

type itemperature interface {
	execute(b *models.Beer)
	setNext(itemperature)
}
