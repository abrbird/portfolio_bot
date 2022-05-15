package service_impl

import (
	"context"
	"gitlab.ozon.dev/zBlur/homework-2/internal/domain"
	"gitlab.ozon.dev/zBlur/homework-2/internal/repository"
)

type CurrencyService struct{}

func (c CurrencyService) Retrieve(ctx context.Context, code string, repo repository.CurrencyRepository) domain.CurrencyRetrieve {
	return repo.Retrieve(ctx, code)
}
