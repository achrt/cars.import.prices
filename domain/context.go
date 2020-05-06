package domain

import (
	"errors"

	"cars.import.prices/domain/services"
)

type Context interface {
	Session() *Session
	Logger() Logger
	Services() services.Services
}

func ValidateContext(c Context) error {
	if c == nil {
		return errors.New("Validate context error: Context should not be nil")
	}

	if c.Logger() == nil {
		return errors.New("Validate context error: Logger should not be nil")
	}

	return nil
}
