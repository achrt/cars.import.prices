package replace_price

import (
	"errors"
	"fmt"

	"cars.import.prices/domain/model"

	"cars.import.prices/domain"
)

type Request struct {
	CarHook *model.CarHook
}

type Response struct {
	CarHook *model.CarHook
}

func Run(c domain.Context, req *Request) (*Response, error) {

	if req.CarHook.IsNew == false {
		return &Response{req.CarHook}, nil
	}

	if req.CarHook.MarketingComplectationID == "" {
		return nil, errors.New("Marketing complectation code should not be empty")
	}

	prices, err := c.Services().GetPricesByMarkId(req.CarHook.MarketingComplectationID, c.Logger())
	if err != nil {
		return nil, err
	}

	if prices == nil {
		return nil, errors.New(fmt.Sprintf("Can not find prices for makreting complectation code %s.", req.CarHook.MarketingComplectationID))
	}

	pt := "price_retail"
	priceType, err := c.Services().GetPriceTypeByCode(pt, c.Logger())

	if err != nil {
		return nil, err
	}

	if priceType == nil {
		return nil, errors.New(fmt.Sprintf("Can not find price type by code %s.", pt))
	}

	retail := 0

	for _, price := range prices.Attributes.Values {
		if price.PriceTypeId == priceType.Id {
			retail = int(price.Value)
		}
	}

	if retail == 0 {
		return nil, errors.New(fmt.Sprintf("Can not find price typed of %s for makreting complectation code %s.", pt, req.CarHook.MarketingComplectationID))
	}

	req.CarHook.Price = retail
	return &Response{req.CarHook}, nil
}
