package replace_price

import (
	"errors"
	"fmt"
	"strconv"

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

	if req.CarHook.MarketingComplectationID == "" || req.CarHook.BodyColorID == "" || req.CarHook.RestylingID == "" {
		return nil, helpers.EmptyData()
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

	if _, ok := helpers.ColorPrices[req.CarHook.RestylingID]; !ok {
		return nil, helpers.RestylingColorsNotFound(req.CarHook.RestylingID)
	}

	if _, ok := helpers.ColorPrices[req.CarHook.RestylingID][req.CarHook.BodyColorID]; !ok {
		return nil, helpers.ColorPriceNotFound(req.CarHook.RestylingID, req.CarHook.BodyColorID)
	}

	colorPrice, err := strconv.Atoi(helpers.ColorPrices[req.CarHook.RestylingID][req.CarHook.BodyColorID]["price"])

	if err != nil {
		return nil, err
	}

	req.CarHook.Price = retail + colorPrice
	return &Response{req.CarHook}, nil
}
