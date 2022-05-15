package service

import (
	"context"
	"gitlab.ozon.dev/zBlur/homework-2/internal/domain"
	"gitlab.ozon.dev/zBlur/homework-2/internal/repository"
)

type CurrencyService interface {
	Retrieve(ctx context.Context, code string, repo repository.CurrencyRepository) domain.CurrencyRetrieve
}
