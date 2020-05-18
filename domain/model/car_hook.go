package model

type CarHook struct {
	VIN                      string
	StockImagesID            []string
	DealersImagesID          []string
	RawTextOptions           []string
	Tags                     []string
	Options                  []string
	AdditionalOptions        []string
	Price                    int
	PriceWithDiscounts       int
	MarketingComplectationID string
	BodyColorID              string
	RestylingID              string
	IsNew                    bool
	Year                     int
}
