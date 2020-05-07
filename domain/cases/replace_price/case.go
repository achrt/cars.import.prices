package replace_price

import (
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
	prices, err := c.Services().GetPricesByMarkId(req.CarHook.MarketingComplectationID, c.Logger())
	if err != nil {
		return nil, err
	}

	if prices != nil {
		fmt.Println(prices.Attributes.Values)
		fmt.Println(prices.Attributes.Status)
	}

	req.CarHook.Price = 0
	return &Response{req.CarHook}, nil
}
