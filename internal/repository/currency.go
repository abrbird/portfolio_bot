package repository

import (
	"context"
	"github.com/abrbird/portfolio_bot/internal/domain"
)

type CurrencyRepository interface {
	Retrieve(ctx context.Context, currencyCode string) domain.CurrencyRetrieve
}
