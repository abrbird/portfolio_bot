package client

import (
	"gitlab.ozon.dev/zBlur/homework-2/internal/domain"
	"gitlab.ozon.dev/zBlur/homework-2/pkg/api"
)

type Client interface {
	GetOrCreateUser(user *domain.User) (*api.User, error)
	GetOrCreatePortfolio(userId int64) (*api.Portfolio, error)
	CreateOrUpdatePortfolioItem(portfolioItemData *domain.PortfolioItemCreate) (*api.PortfolioItem, error)
	DeletePortfolioItem(portfolioItemId int64) (*api.Empty, error)
	GetAvailableMarketItems() ([]*api.MarketItem, error)
	GetMarketItemPrices(marketItemId int64, startTimeStamp int64, endTimeStamp int64, interval int64) ([]*api.MarketPrice, error)
	GetMarketLastPrices(marketItemIds []int64) ([]*api.MarketPrice, error)
}
