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

	fmt.Println(prices)

	req.CarHook.Price = 0
	return &Response{req.CarHook}, nil
}
