package yahoo_finance

import (
	"gitlab.ozon.dev/zBlur/homework-2/internal/domain"
)

type Historical struct {
	PreviousClose      interface{} `json:"previousClose"`
	Symbol             string      `json:"symbol"`
	ChartPreviousClose float64     `json:"chartPreviousClose"`
	Timestamp          []int64     `json:"timestamp"`
	Close              []float64   `json:"close"`
	End                interface{} `json:"end"`
	Start              interface{} `json:"start"`
	DataGranularity    int         `json:"dataGranularity"`
}

func (h *Historical) ToMarketPriceArray(marketItem domain.MarketItem) *[]domain.MarketPrice {
	prices := make([]domain.MarketPrice, len(h.Timestamp))
	for i := 0; i < len(h.Timestamp); i++ {

		prices[i] = domain.MarketPrice{
			MarketItemId: marketItem.Id,
			Price:        h.Close[i],
			Timestamp:    h.Timestamp[i],
		}
	}

	return &prices
}
