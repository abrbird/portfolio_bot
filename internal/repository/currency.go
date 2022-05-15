package repository

import (
	"context"
	"gitlab.ozon.dev/zBlur/homework-2/internal/domain"
)

type CurrencyRepository interface {
	Retrieve(ctx context.Context, currencyCode string) domain.CurrencyRetrieve
}
