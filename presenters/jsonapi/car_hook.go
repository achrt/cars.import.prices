package jsonapi

import (
	"encoding/json"

	"cars.import.prices/domain/model"
)

type CarHook struct {
	Type       string             `json:"type"`
	Attributes *CarHookAttributes `json:"attributes"`
}

type CarHookAttributes struct {
	VIN                      string   `json:"vin"`
	StockImagesID            []string `json:"stock_images_id"`
	DealersImagesID          []string `json:"dealer_images_id"`
	RawTextOptions           []string `json:"raw_text_options"`
	Tags                     []string `json:"tags"`
	Options                  []string `json:"options"`
	AdditionalOptions        []string `json:"additional_options"`
	Price                    int      `json:"price"`
	PriceWithDiscounts       int      `json:"price_with_discounts"`
	MarketingComplectationID string   `json:"marketing_complectation_id"`
	RestylingID              string   `json:"restyling_id"`
	BodyColorID              string   `json:"body_color_id"`
	IsNew                    bool     `json:"is_new"`
	Year                     int      `json:"year"`
}

type CarHookResponse struct {
	Data *CarHook `json:"data"`
}

func MarshalCarHook(carHook *model.CarHook) *CarHookResponse {
	return &CarHookResponse{
		&CarHook{
			Type: "car_hook",
			Attributes: &CarHookAttributes{
				VIN:                      carHook.VIN,
				StockImagesID:            carHook.StockImagesID,
				DealersImagesID:          carHook.StockImagesID,
				RawTextOptions:           carHook.RawTextOptions,
				Tags:                     carHook.Tags,
				Options:                  carHook.Options,
				AdditionalOptions:        carHook.AdditionalOptions,
				Price:                    carHook.Price,
				PriceWithDiscounts:       carHook.PriceWithDiscounts,
				MarketingComplectationID: carHook.MarketingComplectationID,
				RestylingID:              carHook.RestylingID,
				BodyColorID:              carHook.BodyColorID,
				IsNew:                    carHook.IsNew,
				Year:                     carHook.Year,
			},
		},
	}
}

func UnmarshalCarHook(body []byte) (*CarHook, error) {
	result := &CarHookResponse{}
	err := json.Unmarshal(body, result)

	return result.Data, err
}

func (carHook *CarHook) Model() *model.CarHook {
	return &model.CarHook{
		VIN:                      carHook.Attributes.VIN,
		StockImagesID:            carHook.Attributes.StockImagesID,
		DealersImagesID:          carHook.Attributes.StockImagesID,
		RawTextOptions:           carHook.Attributes.RawTextOptions,
		Tags:                     carHook.Attributes.Tags,
		Options:                  carHook.Attributes.Options,
		AdditionalOptions:        carHook.Attributes.AdditionalOptions,
		Price:                    carHook.Attributes.Price,
		PriceWithDiscounts:       carHook.Attributes.PriceWithDiscounts,
		MarketingComplectationID: carHook.Attributes.MarketingComplectationID,
		RestylingID:              carHook.Attributes.RestylingID,
		BodyColorID:              carHook.Attributes.BodyColorID,
		IsNew:                    carHook.Attributes.IsNew,
		Year:                     carHook.Attributes.Year,
	}
}
