package client

import (
	"gitlab.ozon.dev/zBlur/homework-2/internal/domain"
	"gitlab.ozon.dev/zBlur/homework-2/pkg/api"
)

type Client interface {
	GetOrCreateUser(user *domain.User) (*api.User, error)
	GetOrCreatePortfolio(userId int64) (*api.Portfolio, error)
	GetAvailableMarketItems() ([]*api.MarketItem, error)
	GetMarketItemPrices(marketItemId int64, startTimeStamp int64, endTimeStamp int64, interval int64) ([]*api.MarketPrice, error)
}
