package domain

import "cars.import.prices/domain/services"

type Context interface {
	Session() *Session
	Logger() Logger
	Services() Services
}

type Services interface {
	CarsCatalog() CarsCatalog
}

type CarsCatalog interface {
	GetPricesByMarkId(markId string, logger Logger) (*services.Price, error)
	GetPriceTypeByCode(code string, logger Logger) (*services.PriceType, error)
}

type CarsCatalogStruct struct {
}
