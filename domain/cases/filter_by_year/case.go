package filter_by_year

import (
	"errors"
	"fmt"
	"strconv"

	"cars.import.prices/domain"
	"cars.import.prices/domain/model"
)

type Request struct {
	CarHook  *model.CarHook
	YearFrom int
}

type Response struct {
	CarHook *model.CarHook
}

func Run(c domain.Context, req *Request) (*Response, error) {
	if req.YearFrom > req.CarHook.Year {
		return nil, errors.New(fmt.Sprintf("Автомобиль старше %s года.", strconv.Itoa(req.YearFrom)))
	}
	return &Response{req.CarHook}, nil
}
