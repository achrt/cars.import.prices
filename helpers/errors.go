package helpers

import (
	"errors"
	"fmt"
)

func EmptyMarkCode() error {
	return errors.New("Marketing complectation code should not be empty")
}

func PriceOfMarkCodeNotFound(code string) error {
	return errors.New(fmt.Sprintf("Can not find prices for makreting complectation code %s.", code))
}

func PriceTypeNotFound(pt string) error {
	return errors.New(fmt.Sprintf("Can not find price type by code %s.", pt))
}
