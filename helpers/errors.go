package helpers

import (
	"errors"
	"fmt"
)

func EmptyData() error {
	return errors.New("Marketing complectation code, restyling and color ids should not be empty")
}

func PriceOfMarkCodeNotFound(code string) error {
	return errors.New(fmt.Sprintf("Can not find prices for makreting complectation code %s.", code))
}

func PriceTypeNotFound(pt string) error {
	return errors.New(fmt.Sprintf("Can not find a price type by the code %s.", pt))
}

func RestylingColorsNotFound(r string) error {
	return errors.New(fmt.Sprintf("Can not find colors for the restyling id %s.", r))
}

func ColorPriceNotFound(r, c string) error {
	return errors.New(fmt.Sprintf("Can not find color price for the restyling id %s and color %s", r, c))
}
