package client

import (
	"gitlab.ozon.dev/zBlur/homework-2/internal/domain"
	"gitlab.ozon.dev/zBlur/homework-2/pkg/api"
)

type Client interface {
	GetOrCreateUser(user *domain.User) (*api.User, error)
	GetOrCreatePortfolio(userId int64) (*api.Portfolio, error)
}
