package api

import (
	"gitlab.kodixauto.ru/nw/logger"

	"cars.import.prices/domain"
	"cars.import.prices/services"
	"cars.import.prices/services/cars_catalog"
)

type Context struct {
	UserId          string
	ClientIds       []string
	TraceId         string
	InternalRequest string
	Vars            map[string]string
	Body            []byte
	services        *domain.Services
}

type Application struct {
	Connection domain.Connection
}

func (c *Context) Session() *domain.Session {
	return &domain.Session{
		UserId:          c.UserId,
		ClientIds:       c.ClientIds,
		TraceId:         c.TraceId,
		InternalRequest: c.InternalRequest,
	}
}

func (c *Context) Logger() domain.Logger {
	return logger.New(c.TraceId)
}

func (c *Context) Services() domain.Services {
	return &services.Services{
		CarsCatalogService: cars_catalog.New("cars-cars-catalog:8080"),
	}
}
