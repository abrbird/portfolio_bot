package client

import (
	"gitlab.ozon.dev/zBlur/homework-2/internal/data_collector/yahoo_finance"
	"gitlab.ozon.dev/zBlur/homework-2/internal/domain"
)

type Client interface {
	GetHistoricalMap(
		marketItems []domain.MarketItem,
		interval string,
		range_ string,
	) (*map[domain.MarketItem]yahoo_finance.Historical, error)
}
