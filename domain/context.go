package domain

import "cars.import.prices/domain/services"

type Context interface {
	Session() *Session
	Logger() Logger
	Services() CarsCatalog
}

type CarsCatalog interface {
	GetPricesByMarkId(markId string, logger Logger) (*services.Price, error)
}

type CarsCatalogStruct struct {
}
