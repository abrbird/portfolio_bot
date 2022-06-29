package client

import (
	"github.com/abrbird/portfolio_bot/internal/data_collector/yahoo_finance"
	"github.com/abrbird/portfolio_bot/internal/domain"
)

type Client interface {
	GetHistoricalMap(
		marketItems []domain.MarketItem,
		interval string,
		range_ string,
	) (*map[domain.MarketItem]yahoo_finance.Historical, error)
}
