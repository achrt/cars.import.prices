package model

type CarHook struct {
	VIN                string
	StockImagesID      []string
	DealersImagesID    []string
	RawTextOptions     []string
	Tags               []string
	Options            []string
	AdditionalOptions  []string
	Price              int
	PriceWithDiscounts int
}
