package graph

import (
	"bytes"
	"gitlab.ozon.dev/zBlur/homework-2/pkg/api"
)

type Drawer interface {
	MarketItem(
		marketItem *api.MarketItem,
		marketItemPrices []*api.MarketPrice,
		portfolioMarketItem *api.PortfolioItem,
	) (*bytes.Buffer, error)

	PortfolioSummary(
		baseShift float64,
		portfolioItems []*api.PortfolioItem,
		itemsPricesMap map[int64][]*api.MarketPrice,
	) (*bytes.Buffer, error)
}
