package replace_price

import (
	"errors"
	"fmt"

	"cars.import.prices/domain/model"

	"cars.import.prices/domain"
	"cars.import.prices/helpers"
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
		return nil, helpers.EmptyMarkCode()
	}

	prices, err := c.Services().CarsCatalog().GetPricesByMarkId(req.CarHook.MarketingComplectationID, c.Logger())
	if err != nil {
		return nil, err
	}

	if prices == nil {
		return nil, helpers.PriceOfMarkCodeNotFound(req.CarHook.MarketingComplectationID)
	}

	pt := "price_retail"
	priceType, err := c.Services().CarsCatalog().GetPriceTypeByCode(pt, c.Logger())

	if err != nil {
		return nil, err
	}

	if priceType == nil {
		return nil, helpers.PriceTypeNotFound(pt)
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
