package api

import (
	"gitlab.kodixauto.ru/nw/logger"

	"cars.import.prices/domain"
	"cars.import.prices/services"
	domainServices "cars.import.prices/domain/services"
)

type Context struct {
	UserId          string
	ClientIds       []string
	TraceId         string
	InternalRequest string
	Vars            map[string]string
	Body            []byte
	services        *services.Services
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

func (c *Context) Services() domainServices.Services {
	return c.services.WithContext(c)
}
