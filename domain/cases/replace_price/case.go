package replace_price

import (
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
	req.CarHook.Price *= 2
	return &Response{req.CarHook}, nil
}
