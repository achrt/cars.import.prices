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

var price = "price_retail"
var priceMin = "price_retail_min"

func Run(c domain.Context, req *Request) (*Response, error) {

	if req.CarHook.IsNew == false {
		return &Response{req.CarHook}, nil
	}

	err := validate(req)
	if err != nil {
		return nil, err
	}

	prices, err := c.Services().CarsCatalog().GetPricesByMarkId(req.CarHook.MarketingComplectationID, c.Logger())
	if err != nil {
		return nil, err
	}

	if prices == nil {
		return nil, errors.New(fmt.Sprintf("Не найдены цены для марк комплектации %s.", req.CarHook.MarketingComplectationID))
	}

	types := make([]string, 0)
	types = append(types, price, priceMin)

	priceTypes, err := c.Services().CarsCatalog().GetPriceTypeByCode(types, c.Logger())

	if err != nil {
		return nil, err
	}

	priceId := ""
	priceMinId := ""

	for _, t := range priceTypes {
		if t.Code == price {
			priceId = t.Id
		}
		if t.Code == priceMin {
			priceMinId = t.Id
		}
	}

	if priceId == "" || priceMinId == "" {
		return nil, errors.New(fmt.Sprintf("В базе отсутствует цена типа %s или %s", price, priceMin))
	}

	retail := 0
	retailMin := 0

	for _, p := range prices.Attributes.Values {
		if p.PriceTypeId == priceId {
			retail = int(p.Value)
		}
		if p.PriceTypeId == priceMinId {
			retailMin = int(p.Value)
		}
	}

	if retail == 0 || retailMin == 0 {
		return nil, errors.New(fmt.Sprintf("Для марк кода %s отсутствует или равна нулю цена типа %s или %s", req.CarHook.MarketingComplectationID, price, priceMin))
	}

	if _, ok := helpers.ColorPrices[req.CarHook.RestylingID]; !ok {
		return nil, errors.New(fmt.Sprintf("Не найдены наценки на цвет для кода рестайлинга %s", req.CarHook.RestylingID))
	}

	if _, ok := helpers.ColorPrices[req.CarHook.RestylingID][req.CarHook.BodyColorID]; !ok {
		return nil, errors.New(fmt.Sprintf("Не найдена наценка на цвет %s для рестайлинга %s", req.CarHook.BodyColorID, req.CarHook.RestylingID))
	}

	colorPrice, err := strconv.Atoi(helpers.ColorPrices[req.CarHook.RestylingID][req.CarHook.BodyColorID]["price"])

	if err != nil {
		return nil, err
	}

	req.CarHook.Price = retail + colorPrice
	req.CarHook.PriceWithDiscounts = retailMin + colorPrice
	return &Response{req.CarHook}, nil
}

func validate(req *Request) error {
	if req.CarHook.MarketingComplectationID == "" {
		return errors.New("Отсутствует маркетинговая комплектация")
	}

	if req.CarHook.BodyColorID == "" {
		return errors.New("Отсутствует код кузова")
	}

	if req.CarHook.RestylingID == "" {
		return errors.New("Отсутствует код рестайлинга")
	}
	return nil
}
